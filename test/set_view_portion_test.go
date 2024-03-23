// This file is auto-generated

package test

import (
	"testing"
)

func TestSetViewPortion(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with cascadeCreate'
	resp, err = client.NewSetViewPortion("u_id", "i_id", 1).SetCascadeCreate(true).SetAdditionalData(map[string]interface{}{"answer": 42}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with existing item and user'
	resp, err = client.NewSetViewPortion("entity_id", "entity_id", 0).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with nonexisting item id'
	resp, err = client.NewSetViewPortion("entity_id", "nonex_id", 1).Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with nonexisting user id'
	resp, err = client.NewSetViewPortion("nonex_id", "entity_id", 0.5).Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with invalid portion'
	resp, err = client.NewSetViewPortion("entity_id", "entity_id", -2).Send()
	checkErrorResponse(t, err, 400)

	// it 'fails with invalid sessionId'
	resp, err = client.NewSetViewPortion("entity_id", "entity_id", 0.7).SetSessionId("a****").Send()
	checkErrorResponse(t, err, 400)
}
