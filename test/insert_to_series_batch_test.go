// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
)

func TestInsertToSeriesInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail when inserting existing item into existing set'
	// it 'does not fail when cascadeCreate is used'
	// it 'really inserts item to the set'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	resp2, err = client.NewAddItem("new_item").Send()
	if err != nil {
		t.Fatal(err)
	}

	resp2, err = client.NewAddItem("new_item3").Send()
	if err != nil {
		t.Fatal(err)
	}

	var reqs = []requests.Request{
		client.NewInsertToSeries("entity_id", "item", "new_item", 3),
		client.NewInsertToSeries("new_set", "item", "new_item2", 1).SetCascadeCreate(true),
		client.NewInsertToSeries("entity_id", "item", "new_item3", 2),
		client.NewInsertToSeries("entity_id", "item", "new_item3", 2),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 200, batchResponse[2].StatusCode)
	assertEqual(t, 409, batchResponse[3].StatusCode)
}
