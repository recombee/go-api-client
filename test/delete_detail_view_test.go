// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestDeleteDetailView(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	interactionsTestSetUp(t)
	client := createClient(t)

	// it 'does not fail with existing entity id'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteDetailView("user", "item").SetTimestamp(time.Unix(0, 0)).Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteDetailView("user", "item").Send()
	checkErrorResponse(t, err, 404)
}
