// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"testing"
)

func TestUpdateMoreItems(t *testing.T) {
	var resp bindings.UpdateMoreItemsResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'updates more items'
	resp, err = client.NewUpdateMoreItems("'int_property' == 42", map[string]interface{}{"int_property": 77}).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp.ItemIds))
	assertEqual(t, 1, resp.Count)
}
