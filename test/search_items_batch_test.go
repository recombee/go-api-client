// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/bindings"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
)

func TestSearchItemsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	recommsTestSetUp(t)
	// it 'finds "hello"'
	// it 'does not find random string'
	// it 'returnProperties'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewSearchItems("entity_id", "hell", 9),
		client.NewSearchItems("entity_id", "sdhskldf", 9),
		client.NewSearchItems("entity_id", "hell", 9).SetReturnProperties(true),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*bindings.SearchResponse))).Recomms))
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 0, len((*(batchResponse[1].Result.(*bindings.SearchResponse))).Recomms))
	assertEqual(t, 200, batchResponse[2].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[2].Result.(*bindings.SearchResponse))).Recomms))
}
