// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
	"time"
)

func TestListScenarios(t *testing.T) {
	var resp []bindings.Scenario
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists scenarios'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListScenarios().Send()
	if err != nil {
		t.Fatal(err)
	}
}
