// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// GetSegmentation Get existing Segmentation.
type GetSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewGetSegmentation creates GetSegmentation request.
// Get existing Segmentation.
func NewGetSegmentation(client ApiClient, segmentationId string) *GetSegmentation {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &GetSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/segmentations/list/%s", segmentationId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(bindings.Segmentation),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *GetSegmentation) SendWithContext(ctx context.Context) (bindings.Segmentation, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.Segmentation{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.Segmentation)), err
}

// Sends the request to the Recombee API
func (r *GetSegmentation) Send() (bindings.Segmentation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
