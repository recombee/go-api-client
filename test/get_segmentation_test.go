// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
)

func TestGetSegmentation(t *testing.T) {
	var resp bindings.Segmentation
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'gets existing segmentation'
	resp2, err = client.NewCreatePropertyBasedSegmentation("seg1", "items", "str_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewGetSegmentation("seg1").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "seg1", resp.SegmentationId)
}
