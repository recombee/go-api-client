// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/bindings"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
)

func TestAddSearchSynonymInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'adds search synonym'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewAddSearchSynonym("sci-fi", "science fiction").SetOneWay(true),
		client.NewAddSearchSynonym("sci-fi", "science fiction").SetOneWay(true),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 201, batchResponse[0].StatusCode)
	assertEqual(t, "sci-fi", (*(batchResponse[0].Result.(*bindings.SearchSynonym))).Term)
	assertEqual(t, "science fiction", (*(batchResponse[0].Result.(*bindings.SearchSynonym))).Synonym)
	assertEqual(t, 409, batchResponse[1].StatusCode)
}
