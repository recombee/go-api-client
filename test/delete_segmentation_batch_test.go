// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
	"time"
)

func TestDeleteSegmentationInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'deletes segmentation'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	time.Sleep(10 * time.Second)
	resp2, err = client.NewCreatePropertyBasedSegmentation("seg1", "items", "str_property").Send()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(10 * time.Second)
	var reqs = []requests.Request{
		client.NewDeleteSegmentation("seg1"),
		client.NewDeleteSegmentation("seg1"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 404, batchResponse[1].StatusCode)
}
