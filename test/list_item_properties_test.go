// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/recombee/bindings"
	"testing"
	"time"
)

func TestListItemProperties(t *testing.T) {
	var resp []bindings.PropertyInfo
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists properties'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListItemProperties().Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 2, len(resp))
}
