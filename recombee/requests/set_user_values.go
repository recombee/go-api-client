package requests

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// SetUserValues Sets/updates (some) property values of the given user. The properties (columns) must be previously created by [Add user property](https://docs.recombee.com/api.html#add-user-property).
type SetUserValues struct {
	ApiRequest
	client ApiClient
}

// NewSetUserValues creates SetUserValues request.
// Sets/updates (some) property values of the given user. The properties (columns) must be previously created by [Add user property](https://docs.recombee.com/api.html#add-user-property).
func NewSetUserValues(client ApiClient, userId string, values map[string]interface{}) *SetUserValues {

	queryParams := map[string]interface{}{}

	return &SetUserValues{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/users/%s", userId),
			BodyParameters:  values,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * time.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the user should be created if not present in the database.
func (r *SetUserValues) SetCascadeCreate(cascadeCreate bool) *SetUserValues {
	r.BodyParameters["!cascadeCreate"] = cascadeCreate
	return r
}

func (r *SetUserValues) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

func (r *SetUserValues) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
