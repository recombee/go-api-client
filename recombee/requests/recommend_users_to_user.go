// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RecommendUsersToUser Gets users similar to the given user, based on the user's past interactions (purchases, ratings, etc.) and values of properties.
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
// The returned users are sorted by similarity (the first user being the most similar).
type RecommendUsersToUser struct {
	ApiRequest
	client ApiClient
}

// NewRecommendUsersToUser creates RecommendUsersToUser request.
// Gets users similar to the given user, based on the user's past interactions (purchases, ratings, etc.) and values of properties.
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
// The returned users are sorted by similarity (the first user being the most similar).
func NewRecommendUsersToUser(client ApiClient, userId string, count int) *RecommendUsersToUser {

	bodyParameters := map[string]interface{}{
		"count": count,
	}

	queryParams := map[string]interface{}{}

	return &RecommendUsersToUser{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/recomms/users/%s/users/", userId),
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
func (r *RecommendUsersToUser) SetScenario(scenario string) *RecommendUsersToUser {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If the user does not exist in the database, returns a list of non-personalized recommendations and creates the user in the database. This allows, for example, rotations in the following recommendations for that user, as the user will be already known to the system.
func (r *RecommendUsersToUser) SetCascadeCreate(cascadeCreate bool) *RecommendUsersToUser {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the recommended users are returned along with their IDs in a JSON dictionary. The acquired property values can be used to easily display the recommended users.
// Example response:
// ```json
//
//	{
//	  "recommId": "9cb9c55d-50ba-4478-84fd-ab456136156e",
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
func (r *RecommendUsersToUser) SetReturnProperties(returnProperties bool) *RecommendUsersToUser {
	r.BodyParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=country`:
// ```json
//
//	{
//	  "recommId": "b326d82d-5d57-4b45-b362-c9d6f0895855",
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
func (r *RecommendUsersToUser) SetIncludedProperties(includedProperties []string) *RecommendUsersToUser {
	r.BodyParameters["includedProperties"] = includedProperties
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql) expression, which allows you to filter recommended items based on the values of their attributes.
// Filters can also be assigned to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendUsersToUser) SetFilter(filter string) *RecommendUsersToUser {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql) expression, which allows you to boost the recommendation rate of some items based on the values of their attributes.
// Boosters can also be assigned to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendUsersToUser) SetBooster(booster string) *RecommendUsersToUser {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios) in the [Admin UI](https://admin.recombee.com).
func (r *RecommendUsersToUser) SetLogic(logic bindings.Logic) *RecommendUsersToUser {
	r.BodyParameters["logic"] = logic
	return r
}

// SetDiversity sets the diversity parameter.
// **Expert option:** Real number from [0.0, 1.0], which determines how mutually dissimilar the recommended users should be. The default value is 0.0, i.e., no diversification. Value 1.0 means maximal diversification.
func (r *RecommendUsersToUser) SetDiversity(diversity float64) *RecommendUsersToUser {
	r.BodyParameters["diversity"] = diversity
	return r
}

// SetMinRelevance sets the minRelevance parameter.
// **Expert option:** Specifies the threshold of how relevant must the recommended users be. Possible values one of: "low", "medium", "high".
func (r *RecommendUsersToUser) SetMinRelevance(minRelevance string) *RecommendUsersToUser {
	r.BodyParameters["minRelevance"] = minRelevance
	return r
}

// SetRotationRate sets the rotationRate parameter.
// **Expert option:** If your users browse the system in real-time, it may easily happen that you wish to offer them recommendations multiple times. Here comes the question: how much should the recommendations change? Should they remain the same, or should they rotate? Recombee API allows you to control this per request in a backward fashion. You may penalize a user for being recommended in the near past. For the specific user, `rotationRate=1` means maximal rotation, `rotationRate=0` means absolutely no rotation. You may also use, for example, `rotationRate=0.2` for only slight rotation of recommended users.
func (r *RecommendUsersToUser) SetRotationRate(rotationRate float64) *RecommendUsersToUser {
	r.BodyParameters["rotationRate"] = rotationRate
	return r
}

// SetRotationTime sets the rotationTime parameter.
// **Expert option:** Taking *rotationRate* into account, specifies how long it takes for a user to recover from the penalization. For example, `rotationTime=7200.0` means that users recommended less than 2 hours ago are penalized.
func (r *RecommendUsersToUser) SetRotationTime(rotationTime float64) *RecommendUsersToUser {
	r.BodyParameters["rotationTime"] = rotationTime
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *RecommendUsersToUser) SetExpertSettings(expertSettings map[string]interface{}) *RecommendUsersToUser {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *RecommendUsersToUser) SetReturnAbGroup(returnAbGroup bool) *RecommendUsersToUser {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *RecommendUsersToUser) SendWithContext(ctx context.Context) (bindings.RecommendationResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.RecommendationResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.RecommendationResponse)), err
}

// Sends the request to the Recombee API
func (r *RecommendUsersToUser) Send() (bindings.RecommendationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
