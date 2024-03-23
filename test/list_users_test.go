package test

import (
	"github.com/recombee/go-api-client/recombee/bindings"
	"testing"
	"time"
)

func TestListUsers(t *testing.T) {
	var resp []bindings.User
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists entities'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListUsers().Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "entity_id", resp[0].UserId)

	// it 'return properties'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListUsers().SetReturnProperties(true).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp))
}
