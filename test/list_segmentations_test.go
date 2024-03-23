// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/recombee/bindings"
	"testing"
	"time"
)

func TestListSegmentations(t *testing.T) {
	var resp bindings.ListSegmentationsResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists existing segmentations'
	time.Sleep(10 * time.Second)
	resp2, err = client.NewCreatePropertyBasedSegmentation("seg1", "items", "str_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewListSegmentations("items").Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, 1, len(resp.Segmentations))
}
