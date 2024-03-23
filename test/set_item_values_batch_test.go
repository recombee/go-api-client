// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
)

func TestSetItemValuesInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'does not fail with existing entity and property'
	// it 'does not fail with non-ASCII string'
	// it 'sets multiple properties'
	// it 'does not fail with !cascadeCreate'
	// it 'does not fail with cascadeCreate optional parameter'
	// it 'fails with nonexisting entity'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewSetItemValues("entity_id", map[string]interface{}{"int_property": 5}),
		client.NewSetItemValues("entity_id", map[string]interface{}{"str_property": "šřžذ的‎"}),
		client.NewSetItemValues("entity_id", map[string]interface{}{"int_property": 5, "str_property": "test"}),
		client.NewSetItemValues("new_entity", map[string]interface{}{"int_property": 5, "str_property": "test", "!cascadeCreate": true}),
		client.NewSetItemValues("new_entity2", map[string]interface{}{"int_property": 5, "str_property": "test"}).SetCascadeCreate(true),
		client.NewSetItemValues("nonexisting", map[string]interface{}{"int_property": 5}),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 200, batchResponse[2].StatusCode)
	assertEqual(t, 200, batchResponse[3].StatusCode)
	assertEqual(t, 200, batchResponse[4].StatusCode)
	assertEqual(t, 404, batchResponse[5].StatusCode)
}
