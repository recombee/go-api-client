// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestDeleteAllSearchSynonyms(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'deletes all search synonyms'
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteAllSearchSynonyms().Send()
	if err != nil {
		t.Fatal(err)
	}
}
