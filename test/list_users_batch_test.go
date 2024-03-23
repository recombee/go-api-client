package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/bindings"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
	"time"
)

func TestListUsersInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'lists entities'
	// it 'return properties'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewListUsers(),
		client.NewListUsers().SetReturnProperties(true),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, "entity_id", (*batchResponse[0].Result.(*[]bindings.User))[0].UserId)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 1, len((*(batchResponse)[1].Result.(*[]bindings.User))))
}
