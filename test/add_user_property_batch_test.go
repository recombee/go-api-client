// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"testing"
)

func TestAddUserPropertyInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with valid name and type'
	// it 'fails with invalid type'
	// it 'really stores property to the system'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewAddUserProperty("number", "int"),
		client.NewAddUserProperty("str", "string"),
		client.NewAddUserProperty("prop", "integer"),
		client.NewAddUserProperty("number2", "int"),
		client.NewAddUserProperty("number2", "int"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 201, batchResponse[0].StatusCode)
	assertEqual(t, 201, batchResponse[1].StatusCode)
	assertEqual(t, 400, batchResponse[2].StatusCode)
	assertEqual(t, 201, batchResponse[3].StatusCode)
	assertEqual(t, 409, batchResponse[4].StatusCode)
}
