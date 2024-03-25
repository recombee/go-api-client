// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
	"time"
)

func TestDeleteManualReqlSegmentInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'deletes manual ReQL segment'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

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
	var reqs = []requests.Request{
		client.NewDeleteManualReqlSegment("seg1", "first-segment"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
}
