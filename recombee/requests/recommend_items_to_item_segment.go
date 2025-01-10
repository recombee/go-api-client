// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendItemsToItemSegment Recommends Items that are the most relevant to a particular Segment from a context [Segmentation](https://docs.recombee.com/segmentations.html).
// Based on the used Segmentation, this endpoint can be used for example for:
// - Recommending articles related to a particular topic
// - Recommending songs belonging to a particular genre
// - Recommending products produced by a particular brand
// You need to set the used context Segmentation in the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned items are sorted by relevance (the first item being the most relevant).
// It is also possible to use the POST HTTP method (for example, in the case of a very long ReQL filter) — query parameters then become body parameters.
type RecommendItemsToItemSegment struct {
	ApiRequest
	client ApiClient
}

// NewRecommendItemsToItemSegment creates RecommendItemsToItemSegment request.
// Recommends Items that are the most relevant to a particular Segment from a context [Segmentation](https://docs.recombee.com/segmentations.html).
// Based on the used Segmentation, this endpoint can be used for example for:
// - Recommending articles related to a particular topic
// - Recommending songs belonging to a particular genre
// - Recommending products produced by a particular brand
// You need to set the used context Segmentation in the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned items are sorted by relevance (the first item being the most relevant).
// It is also possible to use the POST HTTP method (for example, in the case of a very long ReQL filter) — query parameters then become body parameters.
func NewRecommendItemsToItemSegment(client ApiClient, contextSegmentId string, targetUserId string, count int) *RecommendItemsToItemSegment {

	bodyParameters := map[string]interface{}{
		"contextSegmentId": contextSegmentId,
		"targetUserId":     targetUserId,
		"count":            count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendItemsToItemSegment{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/item-segments/items/"),
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
// You can set various settings to the [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com). You can also see the performance of each scenario in the Admin UI separately, so you can check how well each application performs.
// The AI that optimizes models to get the best results may optimize different scenarios separately or even use different models in each of the scenarios.
func (r *RecommendItemsToItemSegment) SetScenario(scenario string) *RecommendItemsToItemSegment {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If an item of the given *itemId* or user of the given *targetUserId* doesn't exist in the database, it creates the missing entity/entities and returns some (non-personalized) recommendations. This allows, for example, rotations in the following recommendations for the user of the given *targetUserId*, as the user will be already known to the system.
func (r *RecommendItemsToItemSegment) SetCascadeCreate(cascadeCreate bool) *RecommendItemsToItemSegment {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the recommended items are returned along with their IDs in a JSON dictionary. The acquired property values can be used to easily display the recommended items to the user.
// Example response:
// ```
//
//	{
//	  "recommId": "0c6189e7-dc1a-429a-b613-192696309361",
//	  "recomms":
//	    [
//	      {
//	        "id": "tv-178",
//	        "values": {
//	          "description": "4K TV with 3D feature",
//	          "categories":   ["Electronics", "Televisions"],
//	          "price": 342,
//	          "url": "myshop.com/tv-178"
//	        }
//	      },
//	      {
//	        "id": "mixer-42",
//	        "values": {
//	          "description": "Stainless Steel Mixer",
//	          "categories":   ["Home & Kitchen"],
//	          "price": 39,
//	          "url": "myshop.com/mixer-42"
//	        }
//	      }
//	    ],
//	  "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *RecommendItemsToItemSegment) SetReturnProperties(returnProperties bool) *RecommendItemsToItemSegment {
	r.BodyParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=description,price`:
// ```
//
//	{
//	  "recommId": "6842c725-a79f-4537-a02c-f34d668a3f80",
//	  "recomms":
//	    [
//	      {
//	        "id": "tv-178",
//	        "values": {
//	          "description": "4K TV with 3D feature",
//	          "price": 342
//	        }
//	      },
//	      {
//	        "id": "mixer-42",
//	        "values": {
//	          "description": "Stainless Steel Mixer",
//	          "price": 39
//	        }
//	      }
//	    ],
//	  "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *RecommendItemsToItemSegment) SetIncludedProperties(includedProperties []string) *RecommendItemsToItemSegment {
	r.BodyParameters["includedProperties"] = includedProperties
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to filter recommended items based on the values of their attributes.
// Filters can also be assigned to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToItemSegment) SetFilter(filter string) *RecommendItemsToItemSegment {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to boost the recommendation rate of some items based on the values of their attributes.
// Boosters can also be assigned to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToItemSegment) SetBooster(booster string) *RecommendItemsToItemSegment {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics.html) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToItemSegment) SetLogic(logic bindings.Logic) *RecommendItemsToItemSegment {
	r.BodyParameters["logic"] = logic
	return r
}

// SetMinRelevance sets the minRelevance parameter.
// **Expert option** If the *targetUserId* is provided:  Specifies the threshold of how relevant must the recommended items be to the user. Possible values one of: "low", "medium", "high". The default value is "low", meaning that the system attempts to recommend a number of items equal to *count* at any cost. If there is not enough data (such as interactions or item properties), this may even lead to bestseller-based recommendations being appended to reach the full *count*. This behavior may be suppressed by using "medium" or "high" values. In such case, the system only recommends items of at least the requested relevance and may return less than *count* items when there is not enough data to fulfill it.
func (r *RecommendItemsToItemSegment) SetMinRelevance(minRelevance string) *RecommendItemsToItemSegment {
	r.BodyParameters["minRelevance"] = minRelevance
	return r
}

// SetRotationRate sets the rotationRate parameter.
// **Expert option** If the *targetUserId* is provided: If your users browse the system in real-time, it may easily happen that you wish to offer them recommendations multiple times. Here comes the question: how much should the recommendations change? Should they remain the same, or should they rotate? Recombee API allows you to control this per request in a backward fashion. You may penalize an item for being recommended in the near past. For the specific user, `rotationRate=1` means maximal rotation, `rotationRate=0` means absolutely no rotation. You may also use, for example, `rotationRate=0.2` for only slight rotation of recommended items.
func (r *RecommendItemsToItemSegment) SetRotationRate(rotationRate float64) *RecommendItemsToItemSegment {
	r.BodyParameters["rotationRate"] = rotationRate
	return r
}

// SetRotationTime sets the rotationTime parameter.
// **Expert option** If the *targetUserId* is provided: Taking *rotationRate* into account, specifies how long it takes for an item to recover from the penalization. For example, `rotationTime=7200.0` means that items recommended less than 2 hours ago are penalized.
func (r *RecommendItemsToItemSegment) SetRotationTime(rotationTime float64) *RecommendItemsToItemSegment {
	r.BodyParameters["rotationTime"] = rotationTime
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *RecommendItemsToItemSegment) SetExpertSettings(expertSettings map[string]interface{}) *RecommendItemsToItemSegment {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *RecommendItemsToItemSegment) SetReturnAbGroup(returnAbGroup bool) *RecommendItemsToItemSegment {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendItemsToItemSegment) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendItemsToItemSegment) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
