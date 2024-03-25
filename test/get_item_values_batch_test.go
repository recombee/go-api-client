// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
)

func TestGetItemValuesInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'gets values'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewGetItemValues("entity_id"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 42, (*(batchResponse[0].Result.(*map[string]interface{})))["int_property"])
	assertEqual(t, "hello", (*(batchResponse[0].Result.(*map[string]interface{})))["str_property"])
}
