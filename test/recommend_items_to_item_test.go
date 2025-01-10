// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"testing"
)

func TestRecommendItemsToItem(t *testing.T) {
	var resp bindings.RecommendationResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	recommsTestSetUp(t)
	client := createClient(t)

	// it 'recommends'
	resp, err = client.NewRecommendItemsToItem("entity_id", "entity_id", 9).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'recommends to previously nonexisting entity with cascadeCreate'
	resp, err = client.NewRecommendItemsToItem("nonexisting", "entity_id", 9).SetCascadeCreate(true).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'recommends with expert settings'
	resp, err = client.NewRecommendItemsToItem("nonexisting2", "entity_id", 9).SetCascadeCreate(true).SetExpertSettings(map[string]interface{}{}).Send()
	if err != nil {
		t.Fatal(err)
	}
}
