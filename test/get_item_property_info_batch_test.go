// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
)

func TestGetItemPropertyInfoInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with existing properties'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewGetItemPropertyInfo("int_property"),
		client.NewGetItemPropertyInfo("str_property"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, "int", (*(batchResponse[0].Result.(*bindings.PropertyInfo))).Type)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, "string", (*(batchResponse[1].Result.(*bindings.PropertyInfo))).Type)
}
