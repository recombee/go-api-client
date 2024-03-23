// This file is auto-generated

package test

import (
	"testing"
)

func TestSetUserValues(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with existing entity and property'
	resp, err = client.NewSetUserValues("entity_id", map[string]interface{}{"int_property": 5}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with non-ASCII string'
	resp, err = client.NewSetUserValues("entity_id", map[string]interface{}{"str_property": "šřžذ的‎"}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'sets multiple properties'
	resp, err = client.NewSetUserValues("entity_id", map[string]interface{}{"int_property": 5, "str_property": "test"}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with !cascadeCreate'
	resp, err = client.NewSetUserValues("new_entity", map[string]interface{}{"int_property": 5, "str_property": "test", "!cascadeCreate": true}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with cascadeCreate optional parameter'
	resp, err = client.NewSetUserValues("new_entity2", map[string]interface{}{"int_property": 5, "str_property": "test"}).SetCascadeCreate(true).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with nonexisting entity'
	resp, err = client.NewSetUserValues("nonexisting", map[string]interface{}{"int_property": 5}).Send()
	checkErrorResponse(t, err, 404)
}
