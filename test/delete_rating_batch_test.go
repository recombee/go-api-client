// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
	"time"
)

func TestDeleteRatingInBatch(t *testing.T) {
	defaultTestSetUp(t)
	interactionsTestSetUp(t)
	// it 'does not fail with existing entity id'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteRating("user", "item").SetTimestamp(time.Unix(0, 0)),
		client.NewDeleteRating("user", "item"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 404, batchResponse[1].StatusCode)
}
