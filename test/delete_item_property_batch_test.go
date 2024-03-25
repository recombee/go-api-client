// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
	"time"
)

func TestDeleteItemPropertyInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with existing property'
	// it 'fails with invalid property'
	// it 'fails with non-existing property'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteItemProperty("int_property"),
		client.NewDeleteItemProperty("int_property"),
		client.NewDeleteItemProperty("***not_valid$$$"),
		client.NewDeleteItemProperty("not_existing"),
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
