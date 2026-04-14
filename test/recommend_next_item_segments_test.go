// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
)

func TestRecommendNextItemSegments(t *testing.T) {
	var resp bindings.RecommendationResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'rejects request with invalid recommId'
	resp2, err = client.NewRecommendNextItemSegments("invalid_recomm_id", 5).Send()
	checkErrorResponse(t, err, 400)

	// it 'rejects request to recommId which does not return item-segments'
	resp2, err = client.NewRecommendItemsToUser("entity_id", 3).Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewRecommendNextItemSegments(resp2.(bindings.RecommendationResponse).RecommId, 5).Send()
	checkErrorResponse(t, err, 400)
}
