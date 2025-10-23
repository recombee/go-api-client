// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
)

func TestRecommendItemSegmentsToItemInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'rejects request to scenario which is not set up'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewRecommendItemSegmentsToItem("entity_id", "entity_id", 5).SetScenario("scenario1").SetCascadeCreate(true),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 400, batchResponse[0].StatusCode)
}
