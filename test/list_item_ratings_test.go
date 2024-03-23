// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/recombee/bindings"
	"testing"
	"time"
)

func TestListItemRatings(t *testing.T) {
	var resp []bindings.Rating
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	interactionsTestSetUp(t)
	client := createClient(t)

	// it 'lists interactions'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListItemRatings("item").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp))
	assertEqual(t, "item", resp[0].ItemId)
	assertEqual(t, "user", resp[0].UserId)
}
