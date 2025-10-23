// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
)

func TestSearchItems(t *testing.T) {
	var resp bindings.SearchResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	recommsTestSetUp(t)
	client := createClient(t)

	// it 'finds "hello"'
	resp, err = client.NewSearchItems("entity_id", "hell", 9).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp.Recomms))

	// it 'does not find random string'
	resp, err = client.NewSearchItems("entity_id", "sdhskldf", 9).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 0, len(resp.Recomms))

	// it 'returnProperties'
	resp, err = client.NewSearchItems("entity_id", "hell", 9).SetReturnProperties(true).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp.Recomms))
}
