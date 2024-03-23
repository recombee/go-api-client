// This file is auto-generated

package test

import (
	"testing"
)

func TestAddItemProperty(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with valid name and type'
	resp, err = client.NewAddItemProperty("number", "int").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddItemProperty("str", "string").Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with invalid type'
	resp, err = client.NewAddItemProperty("prop", "integer").Send()
	checkErrorResponse(t, err, 400)

	// it 'really stores property to the system'
	resp, err = client.NewAddItemProperty("number2", "int").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddItemProperty("number2", "int").Send()
	checkErrorResponse(t, err, 409)
}
