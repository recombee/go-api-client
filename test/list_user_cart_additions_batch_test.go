// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
	"time"
)

func TestListUserCartAdditionsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	interactionsTestSetUp(t)
	// it 'lists user interactions'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewListUserCartAdditions("user"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*[]bindings.CartAddition)))))
	assertEqual(t, "item", (*(batchResponse[0].Result.(*[]bindings.CartAddition)))[0].ItemId)
	assertEqual(t, "user", (*(batchResponse[0].Result.(*[]bindings.CartAddition)))[0].UserId)
}
