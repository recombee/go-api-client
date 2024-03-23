// This file is auto-generated

package test

import (
	"testing"
)

func TestGetUserValues(t *testing.T) {
	var resp map[string]interface{}
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'gets values'
	resp, err = client.NewGetUserValues("entity_id").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 42, resp["int_property"])
	assertEqual(t, "hello", resp["str_property"])
}
