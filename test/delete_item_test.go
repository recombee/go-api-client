// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestDeleteItem(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with existing entity id'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteItem("entity_id").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteItem("entity_id").Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with invalid entity id'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteItem("***not_valid$$$").Send()
	checkErrorResponse(t, err, 400)

	// it 'fails with non-existing entity'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteItem("valid_id").Send()
	checkErrorResponse(t, err, 404)
}
