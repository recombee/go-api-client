// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestAddRating(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with cascadeCreate'
	resp, err = client.NewAddRating("u_id", "i_id", 1).SetCascadeCreate(true).SetAdditionalData(map[string]interface{}{"answer": 42}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'does not fail with existing item and user'
	resp, err = client.NewAddRating("entity_id", "entity_id", 0).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with nonexisting item id'
	resp, err = client.NewAddRating("entity_id", "nonex_id", -1).Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with nonexisting user id'
	resp, err = client.NewAddRating("nonex_id", "entity_id", 0.5).Send()
	checkErrorResponse(t, err, 404)

	// it 'fails with invalid rating'
	resp, err = client.NewAddRating("entity_id", "entity_id", -2).Send()
	checkErrorResponse(t, err, 400)

	// it 'really stores interaction to the system'
	resp, err = client.NewAddRating("u_id", "i_id", 0.3).SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)).Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddRating("u_id", "i_id", 0.3).SetCascadeCreate(true).SetTimestamp(time.Unix(5, 0)).Send()
	checkErrorResponse(t, err, 409)
}
