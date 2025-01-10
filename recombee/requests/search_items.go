// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// SearchItems Full-text personalized search. The results are based on the provided `searchQuery` and also on the user's past interactions (purchases, ratings, etc.) with the items (items more suitable for the user are preferred in the results).
// All the string and set item properties are indexed by the search engine.
// This endpoint should be used in a search box on your website/app. It can be called multiple times as the user is typing the query in order to get the most viable suggestions based on the current state of the query, or once after submitting the whole query.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this search was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui.html#reported-metrics).
// - Get subsequent search results when the user scrolls down or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api.html#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
type SearchItems struct {
	ApiRequest
	client ApiClient
}

// NewSearchItems creates SearchItems request.
// Full-text personalized search. The results are based on the provided `searchQuery` and also on the user's past interactions (purchases, ratings, etc.) with the items (items more suitable for the user are preferred in the results).
// All the string and set item properties are indexed by the search engine.
// This endpoint should be used in a search box on your website/app. It can be called multiple times as the user is typing the query in order to get the most viable suggestions based on the current state of the query, or once after submitting the whole query.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this search was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui.html#reported-metrics).
// - Get subsequent search results when the user scrolls down or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api.html#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
func NewSearchItems(client ApiClient, userId string, searchQuery string, count int) *SearchItems {

	bodyParameters := map[string]interface{}{
		"searchQuery": searchQuery,
		"count":       count,
	}

	queryParams := map[string]interface{}{}

	return &SearchItems{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/search/users/%s/items/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.SearchResponse),
		},
		client,
	}
}

// SetScenario sets the scenario parameter.
// Scenario defines a particular search field in your user interface.
// You can set various settings to the [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com). You can also see the performance of each scenario in the Admin UI separately, so you can check how well each field performs.
// The AI that optimizes models to get the best results may optimize different scenarios separately, or even use different models in each of the scenarios.
func (r *SearchItems) SetScenario(scenario string) *SearchItems {
	r.BodyParameters["scenario"] = scenario
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If the user does not exist in the database, returns a list of non-personalized search results and creates the user in the database. This allows, for example, rotations in the following recommendations for that user, as the user will be already known to the system.
func (r *SearchItems) SetCascadeCreate(cascadeCreate bool) *SearchItems {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the recommended items are returned along with their IDs in a JSON dictionary. The acquired property values can be used to easily display the recommended items to the user.
// Example response:
// ```
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
//	  "numberNextRecommsCalls": 0
//	}
//
// ```
func (r *SearchItems) SetReturnProperties(returnProperties bool) *SearchItems {
	r.BodyParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=description,price`:
// ```
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
func (r *SearchItems) SetIncludedProperties(includedProperties []string) *SearchItems {
	r.BodyParameters["includedProperties"] = includedProperties
	return r
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to filter recommended items based on the values of their attributes.
// Filters can also be assigned to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *SearchItems) SetFilter(filter string) *SearchItems {
	r.BodyParameters["filter"] = filter
	return r
}

// SetBooster sets the booster parameter.
// Number-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to boost the recommendation rate of some items based on the values of their attributes.
// Boosters can also be assigned to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *SearchItems) SetBooster(booster string) *SearchItems {
	r.BodyParameters["booster"] = booster
	return r
}

// SetLogic sets the logic parameter.
// Logic specifies the particular behavior of the recommendation models. You can pick tailored logic for your domain and use case.
// See [this section](https://docs.recombee.com/recommendation_logics.html) for a list of available logics and other details.
// The difference between `logic` and `scenario` is that `logic` specifies mainly behavior, while `scenario` specifies the place where recommendations are shown to the users.
// Logic can also be set to a [scenario](https://docs.recombee.com/scenarios.html) in the [Admin UI](https://admin.recombee.com).
func (r *SearchItems) SetLogic(logic bindings.Logic) *SearchItems {
	r.BodyParameters["logic"] = logic
	return r
}

// SetExpertSettings sets the expertSettings parameter.
// Dictionary of custom options.
func (r *SearchItems) SetExpertSettings(expertSettings map[string]interface{}) *SearchItems {
	r.BodyParameters["expertSettings"] = expertSettings
	return r
}

// SetReturnAbGroup sets the returnAbGroup parameter.
// If there is a custom AB-testing running, return the name of the group to which the request belongs.
func (r *SearchItems) SetReturnAbGroup(returnAbGroup bool) *SearchItems {
	r.BodyParameters["returnAbGroup"] = returnAbGroup
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *SearchItems) SendWithContext(ctx context.Context) (bindings.SearchResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.SearchResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.SearchResponse)), err
}

// Sends the request to the Recombee API
func (r *SearchItems) Send() (bindings.SearchResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
