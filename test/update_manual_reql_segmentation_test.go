// This file is auto-generated

package test

import (
	"testing"
)

func TestUpdateManualReqlSegmentation(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'updates manual ReQL segmentation'
	resp2, err = client.NewCreateManualReqlSegmentation("seg1", "items").SetTitle("Test Segmentation").SetDescription("For test purposes").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewUpdateManualReqlSegmentation("seg1").SetTitle("New title").SetDescription("Updated").Send()
	if err != nil {
		t.Fatal(err)
	}
}
