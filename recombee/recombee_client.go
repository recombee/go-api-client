package recombee

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	errorspkg "github.com/recombee/go-api-client/v6/recombee/errors"
	requestspkg "github.com/recombee/go-api-client/v6/recombee/requests"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// RecombeeClient is a client for interacting with the Recombee API.
// It holds configuration and state for making requests to the API, including
// authentication and network settings.
type RecombeeClient struct {
	databaseId        string       // The ID of the database to interact with.
	privateToken      string       // Authentication token for API access.
	baseUri           string       // The base URI for API requests.
	useHttpsByDefault bool         // Flag indicating if HTTPS should be used by default for requests.
	httpClient        *http.Client // Custom HTTP client for making requests.
}

// ClientOptions represents configuration options that can be used to customize
// the behavior of a RecombeeClient. These options include network settings and
// preferences for how to connect to the Recombee API.
type ClientOptions struct {
	Region            *string      // Optional region to specify the API endpoint's geographical location. Either region or baseUri must be specified.
	UseHttpsByDefault *bool        // Optional flag to override the default use of HTTPS for requests (default: true)
	BaseUri           *string      // Optional custom base URI. Either region or baseUri must be specified.
	Port              *int         // Optional port number for the API endpoint.
	HttpClient        *http.Client // Optional custom HTTP client for making requests.
}

// NewRecombeeClientWithOptions creates a new instance of RecombeeClient with
// the provided database ID, private token, and additional options. This function
// allows for detailed configuration of the client, including network settings
// and API endpoint customization.
//
// It returns a pointer to the initialized RecombeeClient and an error if the initialization fails.
//
// Parameters:
// databaseId: The identifier for the Recombee database the client will interact with.
// privateToken: The secret token associated with the database
// options: ClientOptions for custom configurations of the client, such as HTTP client settings, API region, and whether to use HTTPS.
func NewRecombeeClientWithOptions(databaseId string, privateToken string, options ClientOptions) (*RecombeeClient, error) {
	useHttpsByDefault := true
	if options.UseHttpsByDefault != nil {
		useHttpsByDefault = *options.UseHttpsByDefault
	}

	httpClient := options.HttpClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	var baseUri string
	if options.BaseUri != nil {
		baseUri = *options.BaseUri
		if options.Port != nil {
			baseUri = baseUri + ":" + strconv.Itoa(*options.Port)
		}
	}

	if options.Region != nil {
		if baseUri != "" {
			return nil, errors.New("You cannot specify both region and baseUri")
		}

		switch *options.Region {
		case "ap-se":
			baseUri = "rapi-ap-se.recombee.com"
		case "ca-east":
			baseUri = "rapi-ca-east.recombee.com"
		case "eu-west":
			baseUri = "rapi-eu-west.recombee.com"
		case "us-west":
			baseUri = "rapi-us-west.recombee.com"
		default:
			return nil, errors.New("Unknown region specified")
		}
	}

	if baseUri == "" {
		return nil, errors.New("You must specify either region or baseUri")
	}

	return &RecombeeClient{databaseId: databaseId, privateToken: privateToken, baseUri: baseUri, useHttpsByDefault: useHttpsByDefault, httpClient: httpClient}, nil
}

// NewRecombeeClient creates a new instance of RecombeeClient with
// the provided database ID, associated private token, and region where the database is located (such as "us-west")
//
// # This function is a convenience wrapper around NewRecombeeClientWithOptions, providing a simpler way to create a client
//
// The function returns a pointer to the initialized RecombeeClient and an error if the initialization fails.
//
// Parameters:
// databaseId: The identifier for the Recombee database the client will interact with.
// privateToken: The secret token associated with the database
// region: The region of the Recombee cluster where the database is located
func NewRecombeeClient(databaseId string, privateToken string, region string) (*RecombeeClient, error) {
	return NewRecombeeClientWithOptions(databaseId, privateToken, ClientOptions{Region: &region})
}

func (c *RecombeeClient) signUrlStr(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	qs := u.Query()
	qs.Add("hmac_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	u.RawQuery = qs.Encode()

	h := hmac.New(sha1.New, []byte(c.privateToken))
	_, err = h.Write([]byte(u.String()))
	if err != nil {
		return "", err
	}

	signature := hex.EncodeToString(h.Sum(nil))
	// there will always be at least one option (hmac_timestamp)
	return u.String() + "&hmac_sign=" + signature, nil
}

func (client *RecombeeClient) SendRequestWithContext(ctx context.Context, request *requestspkg.ApiRequest) error {

	parsedUrl, err := url.Parse(request.Path)
	if err != nil {
		return err
	}

	queryParams := parsedUrl.Query()
	for key, value := range request.QueryParameters {
		switch v := value.(type) {
		case []string:
			queryParams.Set(key, strings.Join(v, ","))
		default:
			queryParams.Set(key, fmt.Sprintf("%v", value))
		}
	}

	parsedUrl.RawQuery = queryParams.Encode()
	urlWithParams := fmt.Sprintf("/%v%v", client.databaseId, parsedUrl.String())
	signedUrl, err := client.signUrlStr(urlWithParams)

	protocol := "https://"
	if (!client.useHttpsByDefault) && (!request.EnsureHttps) {
		protocol = "http://"
	}
	fullUrl := fmt.Sprintf("%s%s%s", protocol, client.baseUri, signedUrl)

	// Encode body parameters to JSON, if present
	var requestBody []byte
	if request.BodyParameters != nil {
		requestBody, err = json.Marshal(request.BodyParameters)
		if err != nil {
			return err
		}
	}

	// Create the HTTP request with context
	httpRequest, err := http.NewRequestWithContext(ctx, request.HttpMethod, fullUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set necessary headers
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("User-Agent", "recombee-go-api-client/6.0.0")

	start := time.Now()
	// Send the request
	response, err := client.httpClient.Do(httpRequest)
	elapsed := time.Since(start)

	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return errorspkg.NewTimeoutError(elapsed, request, err)
		}
		return err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Check response status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		// Handle non-successful status codes as needed
		return errorspkg.NewResponseError(response.StatusCode, string(bodyBytes), request)
	}

	// Decode the response body into the target if not nil
	if request.Target != nil {
		if err := json.Unmarshal(bodyBytes, request.Target); err != nil {
			return err
		}
	}

	return nil
}

// Creating requests with custom constructors

// NewSetItemValues creates SetItemValues request.
// Sets/updates (some) property values of the given item. The properties (columns) must be previously created by [Add item property](https://docs.recombee.com/api#add-item-property).
func (c *RecombeeClient) NewSetItemValues(itemId string, values map[string]interface{}) *requestspkg.SetItemValues {
	return requestspkg.NewSetItemValues(c, itemId, values)
}

// NewSetUserValues creates SetUserValues request.
// Sets/updates (some) property values of the given user. The properties (columns) must be previously created by [Add user property](https://docs.recombee.com/api#add-user-property).
func (c *RecombeeClient) NewSetUserValues(userId string, values map[string]interface{}) *requestspkg.SetUserValues {
	return requestspkg.NewSetUserValues(c, userId, values)
}

// NewBatch creates Batch request.
// Batch processing allows you to submit arbitrary sequence of requests within a single HTTPS request.
// Any type of request from the above documentation may be used in the Batch, and the Batch may combine different types of requests arbitrarily as well.
// Using Batch requests is beneficial for example, when synchronizing the catalog of items or uploading historical interaction data, as sending the data in Batch is considerably faster than sending the individual requests (thanks to optimizations and  reducing network and HTTPS overhead).
func (c *RecombeeClient) NewBatch(requests []requestspkg.Request) *requestspkg.Batch {
	return requestspkg.NewBatch(c, requests)
}
