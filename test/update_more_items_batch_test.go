// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
)

func TestUpdateMoreItemsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'updates more items'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewUpdateMoreItems("'int_property' == 42", map[string]interface{}{"int_property": 77}),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*bindings.UpdateMoreItemsResponse))).ItemIds))
	assertEqual(t, 1, (*(batchResponse[0].Result.(*bindings.UpdateMoreItemsResponse))).Count)
}
