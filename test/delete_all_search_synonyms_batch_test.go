// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
	"time"
)

func TestDeleteAllSearchSynonymsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'deletes all search synonyms'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteAllSearchSynonyms(),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
}
