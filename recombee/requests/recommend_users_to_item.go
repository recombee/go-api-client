// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendUsersToItem Recommends users that are likely to be interested in the given item.
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
// The returned users are sorted by predicted interest in the item (the first user being the most interested).
type RecommendUsersToItem struct {
	ApiRequest
	client ApiClient
}

// NewRecommendUsersToItem creates RecommendUsersToItem request.
// Recommends users that are likely to be interested in the given item.
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
// The returned users are sorted by predicted interest in the item (the first user being the most interested).
func NewRecommendUsersToItem(client ApiClient, itemId string, count int) *RecommendUsersToItem {

	bodyParameters := map[string]interface{}{
		"count": count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendUsersToItem{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/items/%s/users/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  50000 * timepkg.Millisecond,
			Target:          new(bindings.RecommendationResponse),
		},
		client,
	}
}

// SetScenario sets the scenario parameter.
// Scenario defines a particular application of recommendations. It can be, for example, "homepage", "cart", or "emailing".
// You can set various settings to the [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com). You can also see the performance of each scenario in the Admin UI separately, so you can check how well each application performs.
// The AI that optimizes models to get the best results may optimize different scenarios separately or even use different models in each of the scenarios.
func (r *RecommendUsersToItem) SetScenario(scenario string) *RecommendUsersToItem {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If an item of the given *itemId* doesn't exist in the database, it creates the missing item.
func (r *RecommendUsersToItem) SetCascadeCreate(cascadeCreate bool) *RecommendUsersToItem {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the recommended users are returned along with their IDs in a JSON dictionary. The acquired property values can be used to easily display the recommended users.
// Example response:
// ```json
//
//	{
//	  "recommId": "039b71dc-b9cc-4645-a84f-62b841eecfce",
//	  "recomms":
//	    [
//	      {
//	        "id": "user-17",
//	        "values": {
//	          "country": "US",
//	          "sex": "F"
//	        }
//	      },
//	      {
//	        "id": "user-2",
//	        "values": {
//	          "country": "CAN",
//	          "sex": "M"
//	        }
//	      }
//	    ],
//	  "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *RecommendUsersToItem) SetReturnProperties(returnProperties bool) *RecommendUsersToItem {
	r.BodyParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=country`:
// ```json
//
//	{
//	  "recommId": "b2b355dd-972a-4728-9c6b-2dc229db0678",
//	  "recomms":
//	    [
//	      {
//	        "id": "user-17",
//	        "values": {
//	          "country": "US"
//	        }
//	      },
//	      {
//	        "id": "user-2",
//	        "values": {
//	          "country": "CAN"
//	        }
//	      }
//	    ],
//	  "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *RecommendUsersToItem) SetIncludedProperties(includedProperties []string) *RecommendUsersToItem {
	r.BodyParameters["includedProperties"] = includedProperties
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql) expression, which allows you to filter recommended users based on the values of their attributes.
// Filters can also be assigned to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendUsersToItem) SetFilter(filter string) *RecommendUsersToItem {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql) expression, which allows you to boost the recommendation rate of some users based on the values of their attributes.
// Boosters can also be assigned to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendUsersToItem) SetBooster(booster string) *RecommendUsersToItem {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendUsersToItem) SetLogic(logic bindings.Logic) *RecommendUsersToItem {
	r.BodyParameters["logic"] = logic
	return r
}

// SetReqlExpressions sets the reqlExpressions parameter.
// A dictionary of [ReQL](https://docs.recombee.com/reql) expressions that will be executed for each recommended user.
// This can be used to compute additional properties of the recommended users that are not stored in the database.
// The keys are the names of the expressions, and the values are the actual ReQL expressions.
// Example request:
// ```json
//
//	{
//	  "reqlExpressions": {
//	    "isInUsersCity": "context_user[\"city\"] in 'cities'",
//	    "distanceToUser": "earth_distance('location', context_user[\"location\"])",
//	    "isFromSameCompany": "'company' == context_item[\"company\"]"
//	  }
//	}
//
// ```
// Example response:
// ```json
//
//	{
//	  "recommId": "ce52ada4-e4d9-4885-943c-407db2dee837",
//	  "recomms":
//	    [
//	      {
//	        "id": "restaurant-178",
//	        "reqlEvaluations": {
//	          "isInUsersCity": true,
//	          "distanceToUser": 5200.2,
//	          "isFromSameCompany": false
//	        }
//	      },
//	      {
//	        "id": "bar-42",
//	        "reqlEvaluations": {
//	          "isInUsersCity": false,
//	          "distanceToUser": 2516.0,
//	          "isFromSameCompany": true
//	        }
//	      }
//	    ],
//	   "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *RecommendUsersToItem) SetReqlExpressions(reqlExpressions map[string]string) *RecommendUsersToItem {
	r.BodyParameters["reqlExpressions"] = reqlExpressions
	return r
}

// SetDiversity sets the diversity parameter.
// **Expert option:** Real number from [0.0, 1.0], which determines how mutually dissimilar the recommended users should be. The default value is 0.0, i.e., no diversification. Value 1.0 means maximal diversification.
func (r *RecommendUsersToItem) SetDiversity(diversity float64) *RecommendUsersToItem {
	r.BodyParameters["diversity"] = diversity
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *RecommendUsersToItem) SetExpertSettings(expertSettings map[string]interface{}) *RecommendUsersToItem {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *RecommendUsersToItem) SetReturnAbGroup(returnAbGroup bool) *RecommendUsersToItem {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendUsersToItem) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendUsersToItem) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
