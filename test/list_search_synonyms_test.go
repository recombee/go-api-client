// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"testing"
	"time"
)

func TestListSearchSynonyms(t *testing.T) {
	var resp bindings.ListSearchSynonymsResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists search synonyms'
	time.Sleep(10 * time.Second)
	resp2, err = client.NewAddSearchSynonym("sci-fi", "science fiction").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewListSearchSynonyms().Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp.Synonyms))
	time.Sleep(10 * time.Second)
	resp, err = client.NewListSearchSynonyms().SetCount(10).SetOffset(1).Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 0, len(resp.Synonyms))
}
