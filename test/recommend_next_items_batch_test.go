// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
)

func TestRecommendNextItemsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	recommsTestSetUp(t)
	// it 'recommends'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	resp2, err = client.NewRecommendItemsToUser("entity_id", 3).SetReturnProperties(true).Send()
	if err != nil {
		t.Fatal(err)
	}

	var reqs = []requests.Request{
		client.NewRecommendNextItems(resp2.(bindings.RecommendationResponse).RecommId, 3),
		client.NewRecommendNextItems(resp2.(bindings.RecommendationResponse).RecommId, 3),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 3, len((*(batchResponse[0].Result.(*bindings.RecommendationResponse))).Recomms))
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 3, len((*(batchResponse[1].Result.(*bindings.RecommendationResponse))).Recomms))
}
