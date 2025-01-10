// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
	"time"
)

func TestListSearchSynonymsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'lists search synonyms'
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
		client.NewListSearchSynonyms(),
		client.NewListSearchSynonyms().SetCount(10).SetOffset(1),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*bindings.ListSearchSynonymsResponse))).Synonyms))
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 0, len((*(batchResponse[1].Result.(*bindings.ListSearchSynonymsResponse))).Synonyms))
}
