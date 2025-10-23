// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
	"time"
)

func TestDeleteSearchSynonym(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'deletes search synonym'
	time.Sleep(10 * time.Second)
	resp2, err = client.NewAddSearchSynonym("sci-fi", "science fiction").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteSearchSynonym(resp2.(bindings.SearchSynonym).Id).Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteSearchSynonym("a968271b-d666-4dfb-8a30-f459e564ba7b").Send()
	checkErrorResponse(t, err, 404)
}
