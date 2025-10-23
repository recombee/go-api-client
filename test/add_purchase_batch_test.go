// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
	"time"
)

func TestAddPurchaseInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with cascadeCreate'
	// it 'does not fail with existing item and user'
	// it 'does not fail with valid timestamp'
	// it 'fails with nonexisting item id'
	// it 'fails with nonexisting user id'
	// it 'really stores interaction to the system'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewAddPurchase("u_id", "i_id").SetCascadeCreate(true).SetAdditionalData(map[string]interface{}{"answer": 42}),
		client.NewAddPurchase("entity_id", "entity_id"),
		client.NewAddPurchase("entity_id", "entity_id").SetTimestamp(func() time.Time { t, _ := time.Parse(time.RFC3339, "2013-10-29T09:38:41.341Z"); return t }()),
		client.NewAddPurchase("entity_id", "nonex_id"),
		client.NewAddPurchase("nonex_id", "entity_id"),
		client.NewAddPurchase("u_id2", "i_id2").SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)),
		client.NewAddPurchase("u_id2", "i_id2").SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 200, batchResponse[2].StatusCode)
	assertEqual(t, 404, batchResponse[3].StatusCode)
	assertEqual(t, 404, batchResponse[4].StatusCode)
	assertEqual(t, 200, batchResponse[5].StatusCode)
	assertEqual(t, 409, batchResponse[6].StatusCode)
}
