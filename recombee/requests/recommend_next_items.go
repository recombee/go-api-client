// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendNextItems Returns items that shall be shown to a user as next recommendations when the user e.g. scrolls the page down (*infinite scroll*) or goes to the next page.
// It accepts `recommId` of a base recommendation request (e.g., request from the first page) and the number of items that shall be returned (`count`).
// The base request can be one of:
//   - [Recommend items to item](https://docs.recombee.com/api.html#recommend-items-to-item)
//   - [Recommend items to user](https://docs.recombee.com/api.html#recommend-items-to-user)
//   - [Search items](https://docs.recombee.com/api.html#search-items)
//
// All the other parameters are inherited from the base request.
// *Recommend next items* can be called many times for a single `recommId` and each call returns different (previously not recommended) items.
// The number of *Recommend next items* calls performed so far is returned in the `numberNextRecommsCalls` field.
// *Recommend next items* can be requested up to 30 minutes after the base request or a previous *Recommend next items* call.
// For billing purposes, each call to *Recommend next items* is counted as a separate recommendation request.
type RecommendNextItems struct {
	ApiRequest
	client ApiClient
}

// NewRecommendNextItems creates RecommendNextItems request.
// Returns items that shall be shown to a user as next recommendations when the user e.g. scrolls the page down (*infinite scroll*) or goes to the next page.
// It accepts `recommId` of a base recommendation request (e.g., request from the first page) and the number of items that shall be returned (`count`).
// The base request can be one of:
//   - [Recommend items to item](https://docs.recombee.com/api.html#recommend-items-to-item)
//   - [Recommend items to user](https://docs.recombee.com/api.html#recommend-items-to-user)
//   - [Search items](https://docs.recombee.com/api.html#search-items)
//
// All the other parameters are inherited from the base request.
// *Recommend next items* can be called many times for a single `recommId` and each call returns different (previously not recommended) items.
// The number of *Recommend next items* calls performed so far is returned in the `numberNextRecommsCalls` field.
// *Recommend next items* can be requested up to 30 minutes after the base request or a previous *Recommend next items* call.
// For billing purposes, each call to *Recommend next items* is counted as a separate recommendation request.
func NewRecommendNextItems(client ApiClient, recommId string, count int) *RecommendNextItems {

	bodyParameters := map[string]interface{}{
		"count": count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendNextItems{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/next/items/%s", recommId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.RecommendationResponse),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendNextItems) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendNextItems) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
