// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
	"time"
)

func TestDeleteMoreItemsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'deletes more items'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteMoreItems("'int_property' == 42"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*bindings.DeleteMoreItemsResponse))).ItemIds))
	assertEqual(t, 1, (*(batchResponse[0].Result.(*bindings.DeleteMoreItemsResponse))).Count)
}
