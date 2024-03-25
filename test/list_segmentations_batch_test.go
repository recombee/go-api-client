// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
	"time"
)

func TestListSegmentationsInBatch(t *testing.T) {
	defaultTestSetUp(t)
	// it 'lists existing segmentations'
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
		client.NewListSegmentations("items"),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 1, len((*(batchResponse[0].Result.(*bindings.ListSegmentationsResponse))).Segmentations))
}
