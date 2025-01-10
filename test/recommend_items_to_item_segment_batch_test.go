// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
)

func TestRecommendItemsToItemSegmentInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'rejects request to scenario which is not set up'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewRecommendItemsToItemSegment("segment_id", "entity_id", 5).SetScenario("scenario1").SetCascadeCreate(true),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 400, batchResponse[0].StatusCode)
}
