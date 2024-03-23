// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
)

func TestSetViewPortionInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with cascadeCreate'
	// it 'does not fail with existing item and user'
	// it 'fails with nonexisting item id'
	// it 'fails with nonexisting user id'
	// it 'fails with invalid portion'
	// it 'fails with invalid sessionId'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewSetViewPortion("u_id", "i_id", 1).SetCascadeCreate(true).SetAdditionalData(map[string]interface{}{"answer": 42}),
		client.NewSetViewPortion("entity_id", "entity_id", 0),
		client.NewSetViewPortion("entity_id", "nonex_id", 1),
		client.NewSetViewPortion("nonex_id", "entity_id", 0.5),
		client.NewSetViewPortion("entity_id", "entity_id", -2),
		client.NewSetViewPortion("entity_id", "entity_id", 0.7).SetSessionId("a****"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 404, batchResponse[2].StatusCode)
	assertEqual(t, 404, batchResponse[3].StatusCode)
	assertEqual(t, 400, batchResponse[4].StatusCode)
	assertEqual(t, 400, batchResponse[5].StatusCode)
}
