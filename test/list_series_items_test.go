// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
	"time"
)

func TestListSeriesItems(t *testing.T) {
	var resp []bindings.SeriesItem
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists set items'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListSeriesItems("entity_id").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp))
	assertEqual(t, "entity_id", resp[0].ItemId)
	assertEqual(t, "item", resp[0].ItemType)
}
