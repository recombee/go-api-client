// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
	"time"
)

func TestAddRatingInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with cascadeCreate'
	// it 'does not fail with existing item and user'
	// it 'fails with nonexisting item id'
	// it 'fails with nonexisting user id'
	// it 'fails with invalid rating'
	// it 'really stores interaction to the system'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewAddRating("u_id", "i_id", 1).SetCascadeCreate(true).SetAdditionalData(map[string]interface{}{"answer": 42}),
		client.NewAddRating("entity_id", "entity_id", 0),
		client.NewAddRating("entity_id", "nonex_id", -1),
		client.NewAddRating("nonex_id", "entity_id", 0.5),
		client.NewAddRating("entity_id", "entity_id", -2),
		client.NewAddRating("u_id", "i_id", 0.3).SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)),
		client.NewAddRating("u_id", "i_id", 0.3).SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)),
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
	assertEqual(t, 200, batchResponse[5].StatusCode)
	assertEqual(t, 409, batchResponse[6].StatusCode)
}
