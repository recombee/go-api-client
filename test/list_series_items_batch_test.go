// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
	"time"
)

func TestListSeriesItemsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'lists set items'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewListSeriesItems("entity_id"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*[]bindings.SeriesItem)))))
	assertEqual(t, "entity_id", (*(batchResponse[0].Result.(*[]bindings.SeriesItem)))[0].ItemId)
	assertEqual(t, "item", (*(batchResponse[0].Result.(*[]bindings.SeriesItem)))[0].ItemType)
}
