// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
	"time"
)

func TestDeleteSeriesInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with existing entity id'
	// it 'fails with invalid entity id'
	// it 'fails with non-existing entity'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteSeries("entity_id"),
		client.NewDeleteSeries("entity_id"),
		client.NewDeleteSeries("***not_valid$$$"),
		client.NewDeleteSeries("valid_id"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 404, batchResponse[1].StatusCode)
	assertEqual(t, 400, batchResponse[2].StatusCode)
	assertEqual(t, 404, batchResponse[3].StatusCode)
}
