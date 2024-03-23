// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestAddDetailView(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with cascadeCreate'
	resp, err = client.NewAddDetailView("u_id", "i_id").SetCascadeCreate(true).SetAdditionalData(map[string]interface{}{"answer": 42}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with existing item and user'
	resp, err = client.NewAddDetailView("entity_id", "entity_id").Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with valid timestamp'
	resp, err = client.NewAddDetailView("entity_id", "entity_id").SetTimestamp(func() time.Time { t, _ := time.Parse(time.RFC3339, "2013-10-29T09:38:41.341Z"); return t }()).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with nonexisting item id'
	resp, err = client.NewAddDetailView("entity_id", "nonex_id").Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with nonexisting user id'
	resp, err = client.NewAddDetailView("nonex_id", "entity_id").Send()
	checkErrorResponse(t, err, 404)

	// it 'really stores interaction to the system'
	resp, err = client.NewAddDetailView("u_id2", "i_id2").SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)).Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddDetailView("u_id2", "i_id2").SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)).Send()
	checkErrorResponse(t, err, 409)
}
