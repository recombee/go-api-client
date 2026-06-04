// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
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

	// it 'does not fail with valid property role and metadata'
	resp, err = client.NewAddItemProperty("title", "string").SetRole(bindings.PropertyRole{Name: "title"}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with duplicate property role or invalid metadata'
	resp, err = client.NewAddItemProperty("str4", "string").SetRole(bindings.PropertyRole{Name: "summary"}).Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddItemProperty("str4", "string").SetRole(bindings.PropertyRole{Name: "summary"}).Send()
	checkErrorResponse(t, err, 409)
	resp, err = client.NewAddItemProperty("str5", "string").SetRole(bindings.PropertyRole{Name: "titl"}).Send()
	checkErrorResponse(t, err, 404)
}
