// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/recombee/bindings"
	"testing"
)

func TestRecommendNextItems(t *testing.T) {
	var resp bindings.RecommendationResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	recommsTestSetUp(t)
	client := createClient(t)

	// it 'recommends'
	resp2, err = client.NewRecommendItemsToUser("entity_id", 3).SetReturnProperties(true).Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewRecommendNextItems(resp2.(bindings.RecommendationResponse).RecommId, 3).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 3, len(resp.Recomms))
	resp, err = client.NewRecommendNextItems(resp2.(bindings.RecommendationResponse).RecommId, 3).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 3, len(resp.Recomms))
}
