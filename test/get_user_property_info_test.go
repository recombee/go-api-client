// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
)

func TestGetUserPropertyInfo(t *testing.T) {
	var resp bindings.PropertyInfo
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with existing properties'
	resp, err = client.NewGetUserPropertyInfo("int_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "int", resp.Type)
	resp, err = client.NewGetUserPropertyInfo("str_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "string", resp.Type)
}
