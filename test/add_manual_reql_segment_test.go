// This file is auto-generated

package test

import (
	"testing"
)

func TestAddManualReqlSegment(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'adds manual ReQL segment'
	resp2, err = client.NewCreateManualReqlSegmentation("seg1", "items").SetTitle("Test Segmentation").SetDescription("For test purposes").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewAddManualReqlSegment("seg1", "first-segment", "'str_property' != null").SetTitle("First Segment").Send()
	if err != nil {
		t.Fatal(err)
	}
}
