// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
)

func TestUpdatePropertyBasedSegmentationInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'updates property based segmentation'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	resp2, err = client.NewCreatePropertyBasedSegmentation("seg1", "items", "str_property").Send()
	if err != nil {
		t.Fatal(err)
	}

	var reqs = []requests.Request{
		client.NewUpdatePropertyBasedSegmentation("seg1").SetTitle("New title").SetPropertyName("str_property").SetDescription("Updated"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
}
