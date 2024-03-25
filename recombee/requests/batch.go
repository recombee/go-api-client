package requests

import (
	"context"
	"encoding/json"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"net/http"
	"time"
)

// Batch allows you to submit arbitrary sequence of requests within a single HTTPS request.
// Any type of request may be used in the Batch, and the Batch may combine different types of requests arbitrarily as well.
// Using Batch requests is beneficial for example, when synchronizing the catalog of items or uploading historical interaction data, as sending the data in Batch is considerably faster than sending the individual requests (thanks to optimizations and  reducing network and HTTPS overhead).
type Batch struct {
	ApiRequest
	MaxRequestsPerBatchCall int
	requests                []Request
	client                  ApiClient
}

type rawBatchResponse struct {
	Code int             `json:"code,omitempty"`
	Json json.RawMessage `json:"json,omitempty"`
}

// NewBatch creates Batch request.
// Batch processing allows you to submit arbitrary sequence of requests within a single HTTPS request.
// Any type of request from the above documentation may be used in the Batch, and the Batch may combine different types of requests arbitrarily as well.
// Using Batch requests is beneficial for example, when synchronizing the catalog of items or uploading historical interaction data, as sending the data in Batch is considerably faster than sending the individual requests (thanks to optimizations and  reducing network and HTTPS overhead).
func NewBatch(client ApiClient, requests []Request) *Batch {
	totalTimeout := time.Second
	for _, req := range requests {
		totalTimeout += req.getApiRequest().DefaultTimeout
	}

	responses := make([]rawBatchResponse, len(requests))

	return &Batch{
		ApiRequest: ApiRequest{
			HttpMethod:     http.MethodPost,
			Path:           "/batch/",
			BodyParameters: prepareBatchRequestBody(requests),
			DefaultTimeout: totalTimeout,
			EnsureHttps:    true,
			Target:         &responses,
		},
		MaxRequestsPerBatchCall: 10000,
		requests:                requests,
		client:                  client,
	}
}

func prepareBatchRequestBody(requests []Request) map[string]interface{} {
	requestBodies := make([]map[string]interface{}, len(requests))
	for i, reqProvider := range requests {
		req := reqProvider.getApiRequest()
		requestBody := map[string]interface{}{
			"path":   req.Path,
			"method": req.HttpMethod,
			"params": mergeMaps(req.QueryParameters, req.BodyParameters),
		}
		requestBodies[i] = requestBody
	}

	body := map[string]interface{}{
		"requests": requestBodies,
	}

	return body
}

func mergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(m1)+len(m2))
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

// SetDistinctRecomms sets the distinctRecomms parameter.
// Makes all the recommended items for a certain user distinct among multiple recommendation requests in the batch.
func (r *Batch) SetDistinctRecomms(distinctRecomms bool) *Batch {
	r.BodyParameters["distinctRecomms"] = distinctRecomms
	return r
}

// SendWithContext sends the request to the Recombee API using the specified Context
func (r Batch) SendWithContext(ctx context.Context) ([]bindings.BatchResponse, error) {
	var allResults []bindings.BatchResponse

	for start := 0; start < len(r.requests); start += r.MaxRequestsPerBatchCall {
		end := start + r.MaxRequestsPerBatchCall
		if end > len(r.requests) {
			end = len(r.requests)
		}

		smallBatch := NewBatch(r.client, r.requests[start:end])
		err := r.client.SendRequestWithContext(ctx, &smallBatch.ApiRequest)
		if err != nil {
			return nil, err
		}

		rawResponses := *smallBatch.ApiRequest.Target.(*[]rawBatchResponse)
		results := make([]bindings.BatchResponse, len(rawResponses))

		for i, rawResponse := range rawResponses {
			results[i].StatusCode = rawResponse.Code
			if 200 <= rawResponse.Code && rawResponse.Code < 300 {
				err := json.Unmarshal(rawResponse.Json, smallBatch.requests[i].getApiRequest().Target)
				if err != nil {
					return nil, err
				}
				results[i].Result = smallBatch.requests[i].getApiRequest().Target
			} else {
				err := json.Unmarshal(rawResponse.Json, &results[i].Error)
				if err != nil {
					return nil, err
				}
			}
		}

		allResults = append(allResults, results...)
	}

	return allResults, nil
}

// Sends the request to the Recombee API
func (r Batch) Send() ([]bindings.BatchResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
