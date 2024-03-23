// This file is auto-generated

package test

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with valid entity id'
	resp, err = client.NewAddUser("valid_id").Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with invalid entity id'
	resp, err = client.NewAddUser("***not_valid$$$").Send()
	checkErrorResponse(t, err, 400)

	// it 'really stores entity to the system'
	resp, err = client.NewAddUser("valid_id2").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddUser("valid_id2").Send()
	checkErrorResponse(t, err, 409)
}
