// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
)

func TestAddSeriesInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with valid entity id'
	// it 'fails with invalid entity id'
	// it 'really stores entity to the system'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewAddSeries("valid_id"),
		client.NewAddSeries("***not_valid$$$"),
		client.NewAddSeries("valid_id2"),
		client.NewAddSeries("valid_id2"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 201, batchResponse[0].StatusCode)
	assertEqual(t, 400, batchResponse[1].StatusCode)
	assertEqual(t, 201, batchResponse[2].StatusCode)
	assertEqual(t, 409, batchResponse[3].StatusCode)
}
