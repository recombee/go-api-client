// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendNextItemSegments Returns Item segments that shall be shown to a user as next recommendations when the user e.g. scrolls the page down (*infinite scroll*) or goes to the next page.
// It accepts `recommId` of a base recommendation request (e.g., request from the first page) and the number of segments that shall be returned (`count`).
// The base request can be one of:
//   - [Recommend Item Segments to Item](https://docs.recombee.com/api#recommend-item-segments-to-item)
//   - [Recommend Item Segments to User](https://docs.recombee.com/api#recommend-item-segments-to-user)
//   - [Recommend Item Segments to Item Segment](https://docs.recombee.com/api#recommend-item-segments-to-item-segment)
//   - [Search Item Segments](https://docs.recombee.com/api#search-item-segments)
//
// All the other parameters are inherited from the base request.
// *Recommend next Item segments* can be called many times for a single `recommId` and each call returns different (previously not recommended) segments.
// The number of *Recommend next Item segments* calls performed so far is returned in the `numberNextRecommsCalls` field.
// *Recommend next Item segments* can be requested up to 30 minutes after the base request or a previous *Recommend next Item segments* call.
// For billing purposes, each call to *Recommend next Item segments* is counted as a separate recommendation request.
type RecommendNextItemSegments struct {
	ApiRequest
	client ApiClient
}

// NewRecommendNextItemSegments creates RecommendNextItemSegments request.
// Returns Item segments that shall be shown to a user as next recommendations when the user e.g. scrolls the page down (*infinite scroll*) or goes to the next page.
// It accepts `recommId` of a base recommendation request (e.g., request from the first page) and the number of segments that shall be returned (`count`).
// The base request can be one of:
//   - [Recommend Item Segments to Item](https://docs.recombee.com/api#recommend-item-segments-to-item)
//   - [Recommend Item Segments to User](https://docs.recombee.com/api#recommend-item-segments-to-user)
//   - [Recommend Item Segments to Item Segment](https://docs.recombee.com/api#recommend-item-segments-to-item-segment)
//   - [Search Item Segments](https://docs.recombee.com/api#search-item-segments)
//
// All the other parameters are inherited from the base request.
// *Recommend next Item segments* can be called many times for a single `recommId` and each call returns different (previously not recommended) segments.
// The number of *Recommend next Item segments* calls performed so far is returned in the `numberNextRecommsCalls` field.
// *Recommend next Item segments* can be requested up to 30 minutes after the base request or a previous *Recommend next Item segments* call.
// For billing purposes, each call to *Recommend next Item segments* is counted as a separate recommendation request.
func NewRecommendNextItemSegments(client ApiClient, recommId string, count int) *RecommendNextItemSegments {

	bodyParameters := map[string]interface{}{
		"count": count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendNextItemSegments{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/next/item-segments/%s", recommId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.RecommendationResponse),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendNextItemSegments) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendNextItemSegments) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
