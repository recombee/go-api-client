// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
)

func TestRecommendNextItemSegmentsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'rejects request with invalid recommId'
	// it 'rejects request to recommId which does not return item-segments'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	resp2, err = client.NewRecommendNextItemSegments("invalid_recomm_id", 5).Send()
	checkErrorResponse(t, err, 400)

	resp2, err = client.NewRecommendItemsToUser("entity_id", 3).Send()
	if err != nil {
		t.Fatal(err)
	}

	var reqs = []requests.Request{
		client.NewRecommendNextItemSegments(resp2.(bindings.RecommendationResponse).RecommId, 5),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 400, batchResponse[0].StatusCode)
}
