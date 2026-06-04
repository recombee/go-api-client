// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendNextItemSegments Returns [Item Segments](https://docs.recombee.com/segmentations) to be shown as the next recommendations when a user scrolls (e.g., within a carousel or feed of Item Segments such as brands, artists, topics, or categories).
// The request requires the `recommId` of a base recommendation request and the number of Segments to return (`count`).
// The base request can be one of:
//   - [Recommend Item Segments to Item](https://docs.recombee.com/api#recommend-item-segments-to-item)
//   - [Recommend Item Segments to User](https://docs.recombee.com/api#recommend-item-segments-to-user)
//   - [Recommend Item Segments to Item Segment](https://docs.recombee.com/api#recommend-item-segments-to-item-segment)
//   - [Search Item Segments](https://docs.recombee.com/api#search-item-segments)
//
// All other parameters are inherited from the base request associated with the provided `recommId`.
// This endpoint can be called multiple times for a single `recommId`. Each call returns different Item Segments that have not been recommended in previous calls.
// The number of calls made so far is returned in the `numberNextRecommsCalls` field.
// Requests can be made up to 30 minutes after the base request or the most recent Recommend Next Item Segments call.
// For billing purposes, each call to this endpoint is counted as a separate recommendation request.
type RecommendNextItemSegments struct {
	ApiRequest
	client ApiClient
}

// NewRecommendNextItemSegments creates RecommendNextItemSegments request.
// Returns [Item Segments](https://docs.recombee.com/segmentations) to be shown as the next recommendations when a user scrolls (e.g., within a carousel or feed of Item Segments such as brands, artists, topics, or categories).
// The request requires the `recommId` of a base recommendation request and the number of Segments to return (`count`).
// The base request can be one of:
//   - [Recommend Item Segments to Item](https://docs.recombee.com/api#recommend-item-segments-to-item)
//   - [Recommend Item Segments to User](https://docs.recombee.com/api#recommend-item-segments-to-user)
//   - [Recommend Item Segments to Item Segment](https://docs.recombee.com/api#recommend-item-segments-to-item-segment)
//   - [Search Item Segments](https://docs.recombee.com/api#search-item-segments)
//
// All other parameters are inherited from the base request associated with the provided `recommId`.
// This endpoint can be called multiple times for a single `recommId`. Each call returns different Item Segments that have not been recommended in previous calls.
// The number of calls made so far is returned in the `numberNextRecommsCalls` field.
// Requests can be made up to 30 minutes after the base request or the most recent Recommend Next Item Segments call.
// For billing purposes, each call to this endpoint is counted as a separate recommendation request.
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
