// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
	"time"
)

func TestDeleteMoreItems(t *testing.T) {
	var resp bindings.DeleteMoreItemsResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'deletes more items'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteMoreItems("'int_property' == 42").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp.ItemIds))
	assertEqual(t, 1, resp.Count)
}
