// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"testing"
)

func TestAddSearchSynonym(t *testing.T) {
	var resp bindings.SearchSynonym
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'adds search synonym'
	resp, err = client.NewAddSearchSynonym("sci-fi", "science fiction").SetOneWay(true).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "sci-fi", resp.Term)
	assertEqual(t, "science fiction", resp.Synonym)
	resp, err = client.NewAddSearchSynonym("sci-fi", "science fiction").SetOneWay(true).Send()
	checkErrorResponse(t, err, 409)
}
