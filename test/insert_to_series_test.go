// This file is auto-generated

package test

import (
	"testing"
)

func TestInsertToSeries(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail when inserting existing item into existing set'
	resp2, err = client.NewAddItem("new_item").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewInsertToSeries("entity_id", "item", "new_item", 3).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail when cascadeCreate is used'
	resp, err = client.NewInsertToSeries("new_set", "item", "new_item2", 1).SetCascadeCreate(true).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'really inserts item to the set'
	resp2, err = client.NewAddItem("new_item3").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewInsertToSeries("entity_id", "item", "new_item3", 2).Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewInsertToSeries("entity_id", "item", "new_item3", 2).Send()
	checkErrorResponse(t, err, 409)
}
