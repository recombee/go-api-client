// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendItemsToItem Recommends a set of items that are somehow related to one given item, *X*. A typical scenario is when the user *A* is viewing *X*. Then you may display items to the user that he might also be interested in. Recommend items to item request gives you Top-N such items, optionally taking the target user *A* into account.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui.html#reported-metrics).
// - Get subsequent recommended items when the user scrolls down (*infinite scroll*) or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api.html#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
type RecommendItemsToItem struct {
	ApiRequest
	client ApiClient
}

// NewRecommendItemsToItem creates RecommendItemsToItem request.
// Recommends a set of items that are somehow related to one given item, *X*. A typical scenario is when the user *A* is viewing *X*. Then you may display items to the user that he might also be interested in. Recommend items to item request gives you Top-N such items, optionally taking the target user *A* into account.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui.html#reported-metrics).
// - Get subsequent recommended items when the user scrolls down (*infinite scroll*) or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api.html#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
func NewRecommendItemsToItem(client ApiClient, itemId string, targetUserId string, count int) *RecommendItemsToItem {

	bodyParameters := map[string]interface{}{
		"targetUserId": targetUserId,
		"count":        count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendItemsToItem{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/items/%s/items/", itemId),
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
func (r *RecommendItemsToItem) SetScenario(scenario string) *RecommendItemsToItem {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If an item of the given *itemId* or user of the given *targetUserId* doesn't exist in the database, it creates the missing entity/entities and returns some (non-personalized) recommendations. This allows, for example, rotations in the following recommendations for the user of the given *targetUserId*, as the user will be already known to the system.
func (r *RecommendItemsToItem) SetCascadeCreate(cascadeCreate bool) *RecommendItemsToItem {
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
func (r *RecommendItemsToItem) SetReturnProperties(returnProperties bool) *RecommendItemsToItem {
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
func (r *RecommendItemsToItem) SetIncludedProperties(includedProperties []string) *RecommendItemsToItem {
	r.BodyParameters["includedProperties"] = includedProperties
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to filter recommended items based on the values of their attributes.
// Filters can also be assigned to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToItem) SetFilter(filter string) *RecommendItemsToItem {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to boost the recommendation rate of some items based on the values of their attributes.
// Boosters can also be assigned to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToItem) SetBooster(booster string) *RecommendItemsToItem {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics.html) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendItemsToItem) SetLogic(logic bindings.Logic) *RecommendItemsToItem {
	r.BodyParameters["logic"] = logic
	return r
}

// SetUserImpact sets the userImpact parameter.
// **Expert option** If *targetUserId* parameter is present, the recommendations are biased towards the given user. Using *userImpact*, you may control this bias. For an extreme case of `userImpact=0.0`, the interactions made by the user are not taken into account at all (with the exception of history-based blacklisting), for `userImpact=1.0`, you'll get a user-based recommendation. The default value is `0`.
func (r *RecommendItemsToItem) SetUserImpact(userImpact float64) *RecommendItemsToItem {
	r.BodyParameters["userImpact"] = userImpact
	return r
}

// SetDiversity sets the diversity parameter.
// **Expert option** Real number from [0.0, 1.0], which determines how mutually dissimilar the recommended items should be. The default value is 0.0, i.e., no diversification. Value 1.0 means maximal diversification.
func (r *RecommendItemsToItem) SetDiversity(diversity float64) *RecommendItemsToItem {
	r.BodyParameters["diversity"] = diversity
	return r
}

// SetMinRelevance sets the minRelevance parameter.
// **Expert option** If the *targetUserId* is provided:  Specifies the threshold of how relevant must the recommended items be to the user. Possible values one of: "low", "medium", "high". The default value is "low", meaning that the system attempts to recommend a number of items equal to *count* at any cost. If there is not enough data (such as interactions or item properties), this may even lead to bestseller-based recommendations being appended to reach the full *count*. This behavior may be suppressed by using "medium" or "high" values. In such case, the system only recommends items of at least the requested relevance and may return less than *count* items when there is not enough data to fulfill it.
func (r *RecommendItemsToItem) SetMinRelevance(minRelevance string) *RecommendItemsToItem {
	r.BodyParameters["minRelevance"] = minRelevance
	return r
}

// SetRotationRate sets the rotationRate parameter.
// **Expert option** If the *targetUserId* is provided: If your users browse the system in real-time, it may easily happen that you wish to offer them recommendations multiple times. Here comes the question: how much should the recommendations change? Should they remain the same, or should they rotate? Recombee API allows you to control this per request in a backward fashion. You may penalize an item for being recommended in the near past. For the specific user, `rotationRate=1` means maximal rotation, `rotationRate=0` means absolutely no rotation. You may also use, for example, `rotationRate=0.2` for only slight rotation of recommended items.
func (r *RecommendItemsToItem) SetRotationRate(rotationRate float64) *RecommendItemsToItem {
	r.BodyParameters["rotationRate"] = rotationRate
	return r
}

// SetRotationTime sets the rotationTime parameter.
// **Expert option** If the *targetUserId* is provided: Taking *rotationRate* into account, specifies how long it takes for an item to recover from the penalization. For example, `rotationTime=7200.0` means that items recommended less than 2 hours ago are penalized.
func (r *RecommendItemsToItem) SetRotationTime(rotationTime float64) *RecommendItemsToItem {
	r.BodyParameters["rotationTime"] = rotationTime
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *RecommendItemsToItem) SetExpertSettings(expertSettings map[string]interface{}) *RecommendItemsToItem {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *RecommendItemsToItem) SetReturnAbGroup(returnAbGroup bool) *RecommendItemsToItem {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendItemsToItem) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendItemsToItem) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
