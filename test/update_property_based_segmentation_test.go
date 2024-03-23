// This file is auto-generated

package test

import (
	"testing"
)

func TestUpdatePropertyBasedSegmentation(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'updates property based segmentation'
	resp2, err = client.NewCreatePropertyBasedSegmentation("seg1", "items", "str_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewUpdatePropertyBasedSegmentation("seg1").SetTitle("New title").SetPropertyName("str_property").SetDescription("Updated").Send()
	if err != nil {
		t.Fatal(err)
	}
}
