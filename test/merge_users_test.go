// This file is auto-generated

package test

import (
	"testing"
)

func TestMergeUsers(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with existing users'
	resp2, err = client.NewAddUser("target").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewMergeUsers("target", "entity_id").Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'fails with nonexisting user'
	resp, err = client.NewMergeUsers("nonex_id", "entity_id").Send()
	checkErrorResponse(t, err, 404)
}
