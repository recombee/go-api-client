// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
	"time"
)

func TestDeleteSearchSynonymInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'deletes search synonym'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	resp2, err = client.NewAddSearchSynonym("sci-fi", "science fiction").Send()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteSearchSynonym(resp2.(bindings.SearchSynonym).Id),
		client.NewDeleteSearchSynonym("a968271b-d666-4dfb-8a30-f459e564ba7b"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 404, batchResponse[1].StatusCode)
}
