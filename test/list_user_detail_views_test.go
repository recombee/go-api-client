// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"testing"
	"time"
)

func TestListUserDetailViews(t *testing.T) {
	var resp []bindings.DetailView
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	interactionsTestSetUp(t)
	client := createClient(t)

	// it 'lists user interactions'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListUserDetailViews("user").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp))
	assertEqual(t, "item", resp[0].ItemId)
	assertEqual(t, "user", resp[0].UserId)
}
