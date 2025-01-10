// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
)

func TestMergeUsersInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with existing users'
	// it 'fails with nonexisting user'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	resp2, err = client.NewAddUser("target").Send()
	if err != nil {
		t.Fatal(err)
	}

	var reqs = []requests.Request{
		client.NewMergeUsers("target", "entity_id"),
		client.NewMergeUsers("nonex_id", "entity_id"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 404, batchResponse[1].StatusCode)
}
