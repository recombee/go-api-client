// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/recombee/requests"
	"testing"
)

func TestCreateManualReqlSegmentationInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'creates manual ReQL segmentation'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewCreateManualReqlSegmentation("seg1", "items").SetTitle("Test Segmentation").SetDescription("For test purposes"),
		client.NewCreateManualReqlSegmentation("seg1", "items").SetTitle("Test Segmentation").SetDescription("For test purposes"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 201, batchResponse[0].StatusCode)
	assertEqual(t, 409, batchResponse[1].StatusCode)
}
