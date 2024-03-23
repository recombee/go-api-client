// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestDeleteUserProperty(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with existing property'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteUserProperty("int_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteUserProperty("int_property").Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with invalid property'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteUserProperty("***not_valid$$$").Send()
	checkErrorResponse(t, err, 400)

	// it 'fails with non-existing property'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteUserProperty("not_existing").Send()
	checkErrorResponse(t, err, 404)
}
