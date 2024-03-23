// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// SearchItemSegments
// Full-text personalized search that returns Segments from a Segmentation. The results are based on the provided `searchQuery` and also on the user's past interactions (purchases, ratings, etc.).
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Searching within categories or brands
//   - Searching within genres or artists
//
// For example if the user is searching for "iPhone" this endpoint can return "cell phones" category.
// You need to set the used Segmentation the Admin UI in the Scenario settings prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
type SearchItemSegments struct {
	ApiRequest
	client ApiClient
}

// NewSearchItemSegments creates SearchItemSegments request.
// Full-text personalized search that returns Segments from a Segmentation. The results are based on the provided `searchQuery` and also on the user's past interactions (purchases, ratings, etc.).
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Searching within categories or brands
//   - Searching within genres or artists
//
// For example if the user is searching for "iPhone" this endpoint can return "cell phones" category.
// You need to set the used Segmentation the Admin UI in the Scenario settings prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
func NewSearchItemSegments(client ApiClient, userId string, searchQuery string, count int) *SearchItemSegments {

	bodyParameters := map[string]interface{}{
		"searchQuery": searchQuery,
		"count":       count,
	}

	queryParams := map[string]interface{}{}

	return &SearchItemSegments{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/search/users/%s/item-segments/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.SearchResponse),
		},
		client,
	}
}

// SetScenario sets the scenario parameter.
// Scenario defines a particular application of recommendations. It can be, for example, "homepage", "cart", or "emailing".
// You can set various settings to the [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com). You can also see the performance of each scenario in the Admin UI separately, so you can check how well each application performs.
// The AI that optimizes models to get the best results may optimize different scenarios separately or even use different models in each of the scenarios.
func (r *SearchItemSegments) SetScenario(scenario string) *SearchItemSegments {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If the user does not exist in the database, returns a list of non-personalized recommendations and creates the user in the database. This allows, for example, rotations in the following recommendations for that user, as the user will be already known to the system.
func (r *SearchItemSegments) SetCascadeCreate(cascadeCreate bool) *SearchItemSegments {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql.html) expression which allows you to filter recommended segments based on the `segmentationId`.
func (r *SearchItemSegments) SetFilter(filter string) *SearchItemSegments {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql.html) expression which allows you to boost recommendation rate of some segments based on the `segmentationId`.
func (r *SearchItemSegments) SetBooster(booster string) *SearchItemSegments {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics.html) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *SearchItemSegments) SetLogic(logic bindings.Logic) *SearchItemSegments {
	r.BodyParameters["logic"] = logic
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *SearchItemSegments) SetExpertSettings(expertSettings map[string]interface{}) *SearchItemSegments {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *SearchItemSegments) SetReturnAbGroup(returnAbGroup bool) *SearchItemSegments {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *SearchItemSegments) SendWithContext(ctx context.Context) (bindings.SearchResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.SearchResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.SearchResponse)), err
}

// Sends the request to the Recombee API
func (r *SearchItemSegments) Send() (bindings.SearchResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
