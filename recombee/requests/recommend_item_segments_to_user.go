// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendItemSegmentsToUser Recommends the top Segments from a [Segmentation](https://docs.recombee.com/segmentations) for a particular user, based on the user's past interactions.
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Recommending the top categories for the user
//   - Recommending the top genres for the user
//   - Recommending the top brands for the user
//   - Recommending the top artists for the user
//
// You need to set the used Segmentation the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
type RecommendItemSegmentsToUser struct {
	ApiRequest
	client ApiClient
}

// NewRecommendItemSegmentsToUser creates RecommendItemSegmentsToUser request.
// Recommends the top Segments from a [Segmentation](https://docs.recombee.com/segmentations) for a particular user, based on the user's past interactions.
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Recommending the top categories for the user
//   - Recommending the top genres for the user
//   - Recommending the top brands for the user
//   - Recommending the top artists for the user
//
// You need to set the used Segmentation the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
func NewRecommendItemSegmentsToUser(client ApiClient, userId string, count int) *RecommendItemSegmentsToUser {

	bodyParameters := map[string]interface{}{
		"count": count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendItemSegmentsToUser{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/users/%s/item-segments/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.RecommendationResponse),
		},
		client,
	}
}

// SetScenario sets the scenario parameter.
// Scenario defines a particular application of recommendations. It can be, for example, "homepage", "cart", or "emailing".
// You can set various settings to the [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com). You can also see the performance of each scenario in the Admin UI separately, so you can check how well each application performs.
// The AI that optimizes models to get the best results may optimize different scenarios separately or even use different models in each of the scenarios.
func (r *RecommendItemSegmentsToUser) SetScenario(scenario string) *RecommendItemSegmentsToUser {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If the user does not exist in the database, returns a list of non-personalized recommendations and creates the user in the database. This allows, for example, rotations in the following recommendations for that user, as the user will be already known to the system.
func (r *RecommendItemSegmentsToUser) SetCascadeCreate(cascadeCreate bool) *RecommendItemSegmentsToUser {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql) expression which allows you to filter recommended segments based on the `segmentationId`.
func (r *RecommendItemSegmentsToUser) SetFilter(filter string) *RecommendItemSegmentsToUser {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql) expression which allows you to boost recommendation rate of some segments based on the `segmentationId`.
func (r *RecommendItemSegmentsToUser) SetBooster(booster string) *RecommendItemSegmentsToUser {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemSegmentsToUser) SetLogic(logic bindings.Logic) *RecommendItemSegmentsToUser {
	r.BodyParameters["logic"] = logic
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *RecommendItemSegmentsToUser) SetExpertSettings(expertSettings map[string]interface{}) *RecommendItemSegmentsToUser {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *RecommendItemSegmentsToUser) SetReturnAbGroup(returnAbGroup bool) *RecommendItemSegmentsToUser {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendItemSegmentsToUser) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendItemSegmentsToUser) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
