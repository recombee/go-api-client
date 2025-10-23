// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
)

func TestUpdateAutoReqlSegmentationInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'updates auto ReQL segmentation'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	resp2, err = client.NewCreateAutoReqlSegmentation("seg1", "items", "{'str_property'}").SetTitle("Test Segmentation").SetDescription("For test purposes").Send()
	if err != nil {
		t.Fatal(err)
	}

	var reqs = []requests.Request{
		client.NewUpdateAutoReqlSegmentation("seg1").SetTitle("New title").SetExpression("{'str_property' + 'str_property'}").SetDescription("Updated"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
}
