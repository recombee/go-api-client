// This file is auto-generated

package test

import (
	"testing"
)

func TestRemoveFromSeries(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail when removing item that is contained in the set'
	resp, err = client.NewRemoveFromSeries("entity_id", "item", "entity_id").Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails when removing item that is not contained in the set'
	resp, err = client.NewRemoveFromSeries("entity_id", "item", "not_contained").Send()
	checkErrorResponse(t, err, 404)
}
