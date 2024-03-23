// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestDeleteManualReqlSegment(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'deletes manual ReQL segment'
	time.Sleep(10 * time.Second)
	resp2, err = client.NewCreateManualReqlSegmentation("seg1", "items").SetTitle("Test Segmentation").SetDescription("For test purposes").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp2, err = client.NewAddManualReqlSegment("seg1", "first-segment", "'str_property' != null").SetTitle("First Segment").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteManualReqlSegment("seg1", "first-segment").Send()
	if err != nil {
		t.Fatal(err)
	}
}
