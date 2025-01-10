// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
)

func TestRemoveFromSeriesInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail when removing item that is contained in the set'
	// it 'fails when removing item that is not contained in the set'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewRemoveFromSeries("entity_id", "item", "entity_id"),
		client.NewRemoveFromSeries("entity_id", "item", "not_contained"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 404, batchResponse[1].StatusCode)
}
