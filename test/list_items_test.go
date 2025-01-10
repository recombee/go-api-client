package test

import (
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"testing"
	"time"
)

func TestListItems(t *testing.T) {
	var resp []bindings.Item
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists entities'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListItems().Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "entity_id", resp[0].ItemId)

	// it 'return properties'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListItems().SetReturnProperties(true).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp))
}
