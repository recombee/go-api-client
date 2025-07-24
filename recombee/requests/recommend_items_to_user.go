// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendItemsToUser Based on the user's past interactions (purchases, ratings, etc.) with the items, recommends top-N items that are most likely to be of high value for the given user.
// The most typical use cases are recommendations on the homepage, in some "Picked just for you" section, or in email.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui#reported-metrics).
// - Get subsequent recommended items when the user scrolls down (*infinite scroll*) or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
type RecommendItemsToUser struct {
	ApiRequest
	client ApiClient
}

// NewRecommendItemsToUser creates RecommendItemsToUser request.
// Based on the user's past interactions (purchases, ratings, etc.) with the items, recommends top-N items that are most likely to be of high value for the given user.
// The most typical use cases are recommendations on the homepage, in some "Picked just for you" section, or in email.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui#reported-metrics).
// - Get subsequent recommended items when the user scrolls down (*infinite scroll*) or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
func NewRecommendItemsToUser(client ApiClient, userId string, count int) *RecommendItemsToUser {

	bodyParameters := map[string]interface{}{
		"count": count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendItemsToUser{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/users/%s/items/", userId),
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
func (r *RecommendItemsToUser) SetScenario(scenario string) *RecommendItemsToUser {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If the user does not exist in the database, returns a list of non-personalized recommendations and creates the user in the database. This allows, for example, rotations in the following recommendations for that user, as the user will be already known to the system.
func (r *RecommendItemsToUser) SetCascadeCreate(cascadeCreate bool) *RecommendItemsToUser {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the recommended items are returned along with their IDs in a JSON dictionary. The acquired property values can be used to easily display the recommended items to the user.
// Example response:
// ```json
//
//	{
//	  "recommId": "ce52ada4-e4d9-4885-943c-407db2dee837",
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
//	   "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *RecommendItemsToUser) SetReturnProperties(returnProperties bool) *RecommendItemsToUser {
	r.BodyParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=description,price`:
// ```json
//
//	{
//	  "recommId": "a86ee8d5-cd8e-46d1-886c-8b3771d0520b",
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
func (r *RecommendItemsToUser) SetIncludedProperties(includedProperties []string) *RecommendItemsToUser {
	r.BodyParameters["includedProperties"] = includedProperties
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql) expression, which allows you to filter recommended items based on the values of their attributes.
// Filters can also be assigned to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToUser) SetFilter(filter string) *RecommendItemsToUser {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql) expression, which allows you to boost the recommendation rate of some items based on the values of their attributes.
// Boosters can also be assigned to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToUser) SetBooster(booster string) *RecommendItemsToUser {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToUser) SetLogic(logic bindings.Logic) *RecommendItemsToUser {
	r.BodyParameters["logic"] = logic
	return r
}

// SetDiversity sets the diversity parameter.
// **Expert option:** Real number from [0.0, 1.0], which determines how mutually dissimilar the recommended items should be. The default value is 0.0, i.e., no diversification. Value 1.0 means maximal diversification.
func (r *RecommendItemsToUser) SetDiversity(diversity float64) *RecommendItemsToUser {
	r.BodyParameters["diversity"] = diversity
	return r
}

// SetMinRelevance sets the minRelevance parameter.
// **Expert option:** Specifies the threshold of how relevant must the recommended items be to the user. Possible values one of: "low", "medium", "high". The default value is "low", meaning that the system attempts to recommend a number of items equal to *count* at any cost. If there is not enough data (such as interactions or item properties), this may even lead to bestseller-based recommendations to be appended to reach the full *count*. This behavior may be suppressed by using "medium" or "high" values. In such a case, the system only recommends items of at least the requested relevance and may return less than *count* items when there is not enough data to fulfill it.
func (r *RecommendItemsToUser) SetMinRelevance(minRelevance string) *RecommendItemsToUser {
	r.BodyParameters["minRelevance"] = minRelevance
	return r
}

// SetRotationRate sets the rotationRate parameter.
// **Expert option:** If your users browse the system in real-time, it may easily happen that you wish to offer them recommendations multiple times. Here comes the question: how much should the recommendations change? Should they remain the same, or should they rotate? Recombee API allows you to control this per request in a backward fashion. You may penalize an item for being recommended in the near past. For the specific user, `rotationRate=1` means maximal rotation, `rotationRate=0` means absolutely no rotation. You may also use, for example, `rotationRate=0.2` for only slight rotation of recommended items. Default: `0`.
func (r *RecommendItemsToUser) SetRotationRate(rotationRate float64) *RecommendItemsToUser {
	r.BodyParameters["rotationRate"] = rotationRate
	return r
}

// SetRotationTime sets the rotationTime parameter.
// **Expert option:** Taking *rotationRate* into account, specifies how long it takes for an item to recover from the penalization. For example, `rotationTime=7200.0` means that items recommended less than 2 hours ago are penalized. Default: `7200.0`.
func (r *RecommendItemsToUser) SetRotationTime(rotationTime float64) *RecommendItemsToUser {
	r.BodyParameters["rotationTime"] = rotationTime
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *RecommendItemsToUser) SetExpertSettings(expertSettings map[string]interface{}) *RecommendItemsToUser {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *RecommendItemsToUser) SetReturnAbGroup(returnAbGroup bool) *RecommendItemsToUser {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendItemsToUser) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendItemsToUser) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
