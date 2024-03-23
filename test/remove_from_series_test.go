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

	// it 'fails when removing item which have different time'
	resp, err = client.NewRemoveFromSeries("entity_id", "item", "entity_id", 0).Send()
	checkErrorResponse(t, err, 404)

	// it 'does not fail when removing item that is contained in the set'
	resp, err = client.NewRemoveFromSeries("entity_id", "item", "entity_id", 1).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails when removing item that is not contained in the set'
	resp, err = client.NewRemoveFromSeries("entity_id", "item", "not_contained", 1).Send()
	checkErrorResponse(t, err, 404)
}
