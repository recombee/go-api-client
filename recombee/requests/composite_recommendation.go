// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// CompositeRecommendation Composite Recommendation returns both a *source entity* (e.g., an Item or [Item Segment](https://docs.recombee.com/segmentations.html)) and a list of related recommendations in a single response.
// It is ideal for use cases such as personalized homepage sections (*Articles from <category>*), *Because You Watched <movie>*, or *Artists Related to Your Favorite Artist <artist>*.
// See detailed **examples and configuration guidance** in the [Composite Scenarios documentation](https://docs.recombee.com/scenarios#composite-recommendations).
// **Structure**
// The endpoint operates in two stages:
// 1. Recommends the *source* (e.g., an Item Segment or item) to the user.
// 2. Recommends *results* (items or Item Segments) related to that *source*.
// For example, *Articles from <category>* can be decomposed into:
//   - [Recommend Item Segments To User](https://docs.recombee.com/api#recommend-item-segments-to-user) to find the category.
//   - [Recommend Items To Item Segment](https://docs.recombee.com/api#recommend-items-to-item-segment) to recommend articles from that category.
//
// Since the first step uses [Recommend Item Segments To User](https://docs.recombee.com/api#recommend-items-to-user), you must include the `userId` parameter in the *Composite Recommendation* request.
// Each *Composite Recommendation* counts as a single recommendation API request for billing.
// **Stage-specific Parameters**
// Additional parameters can be supplied via [sourceSettings](https://docs.recombee.com/api#composite-recommendation-param-sourceSettings) and [resultSettings](https://docs.recombee.com/api#composite-recommendation-param-resultSettings).
// In the example above:
//   - `sourceSettings` may include any parameter valid for [Recommend Item Segments To User](https://docs.recombee.com/api#recommend-items-to-user) (e.g., `filter`, `booster`).
//   - `resultSettings` may include any parameter valid for [Recommend Items To Item Segment](https://docs.recombee.com/api#recommend-items-to-item-segment).
//
// See [this example](https://docs.recombee.com/api#composite-recommendation-example-setting-parameters-for-individual-stages) for more details.
type CompositeRecommendation struct {
	ApiRequest
	client ApiClient
}

// NewCompositeRecommendation creates CompositeRecommendation request.
// Composite Recommendation returns both a *source entity* (e.g., an Item or [Item Segment](https://docs.recombee.com/segmentations.html)) and a list of related recommendations in a single response.
// It is ideal for use cases such as personalized homepage sections (*Articles from <category>*), *Because You Watched <movie>*, or *Artists Related to Your Favorite Artist <artist>*.
// See detailed **examples and configuration guidance** in the [Composite Scenarios documentation](https://docs.recombee.com/scenarios#composite-recommendations).
// **Structure**
// The endpoint operates in two stages:
// 1. Recommends the *source* (e.g., an Item Segment or item) to the user.
// 2. Recommends *results* (items or Item Segments) related to that *source*.
// For example, *Articles from <category>* can be decomposed into:
//   - [Recommend Item Segments To User](https://docs.recombee.com/api#recommend-item-segments-to-user) to find the category.
//   - [Recommend Items To Item Segment](https://docs.recombee.com/api#recommend-items-to-item-segment) to recommend articles from that category.
//
// Since the first step uses [Recommend Item Segments To User](https://docs.recombee.com/api#recommend-items-to-user), you must include the `userId` parameter in the *Composite Recommendation* request.
// Each *Composite Recommendation* counts as a single recommendation API request for billing.
// **Stage-specific Parameters**
// Additional parameters can be supplied via [sourceSettings](https://docs.recombee.com/api#composite-recommendation-param-sourceSettings) and [resultSettings](https://docs.recombee.com/api#composite-recommendation-param-resultSettings).
// In the example above:
//   - `sourceSettings` may include any parameter valid for [Recommend Item Segments To User](https://docs.recombee.com/api#recommend-items-to-user) (e.g., `filter`, `booster`).
//   - `resultSettings` may include any parameter valid for [Recommend Items To Item Segment](https://docs.recombee.com/api#recommend-items-to-item-segment).
//
// See [this example](https://docs.recombee.com/api#composite-recommendation-example-setting-parameters-for-individual-stages) for more details.
func NewCompositeRecommendation(client ApiClient, scenario string, count int) *CompositeRecommendation {

	bodyParameters := map[string]interface{}{
		"scenario": scenario,
		"count":    count,
	}

	queryParams := map[string]interface{}{}

	return &CompositeRecommendation{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/composite/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.CompositeRecommendationResponse),
		},
		client,
	}
}

// SetItemId sets the itemId parameter.
// ID of the item for which the recommendations are to be generated.
func (r *CompositeRecommendation) SetItemId(itemId string) *CompositeRecommendation {
	r.BodyParameters["itemId"] = itemId
	return r
}

// SetUserId sets the userId parameter.
// ID of the user for which the recommendations are to be generated.
func (r *CompositeRecommendation) SetUserId(userId string) *CompositeRecommendation {
	r.BodyParameters["userId"] = userId
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *CompositeRecommendation) SetLogic(logic bindings.Logic) *CompositeRecommendation {
	r.BodyParameters["logic"] = logic
	return r
}

// SetSegmentId sets the segmentId parameter.
// ID of the segment from `contextSegmentationId` for which the recommendations are to be generated.
func (r *CompositeRecommendation) SetSegmentId(segmentId string) *CompositeRecommendation {
	r.BodyParameters["segmentId"] = segmentId
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If the entity for the source recommendation does not exist in the database, returns a list of non-personalized recommendations and creates the user in the database. This allows, for example, rotations in the following recommendations for that entity, as the entity will be already known to the system.
func (r *CompositeRecommendation) SetCascadeCreate(cascadeCreate bool) *CompositeRecommendation {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetSourceSettings sets the sourceSettings parameter.
// Parameters applied for recommending the *Source* stage. The accepted parameters correspond with the recommendation sub-endpoint used to recommend the *Source*.
func (r *CompositeRecommendation) SetSourceSettings(sourceSettings bindings.CompositeRecommendationStageParameters) *CompositeRecommendation {
	r.BodyParameters["sourceSettings"] = sourceSettings
	return r
}

// SetResultSettings sets the resultSettings parameter.
// Parameters applied for recommending the *Result* stage. The accepted parameters correspond with the recommendation sub-endpoint used to recommend the *Result*.
func (r *CompositeRecommendation) SetResultSettings(resultSettings bindings.CompositeRecommendationStageParameters) *CompositeRecommendation {
	r.BodyParameters["resultSettings"] = resultSettings
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *CompositeRecommendation) SetExpertSettings(expertSettings map[string]interface{}) *CompositeRecommendation {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *CompositeRecommendation) SendWithContext(ctx context.Context) (bindings.CompositeRecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.CompositeRecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.CompositeRecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *CompositeRecommendation) Send() (bindings.CompositeRecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
