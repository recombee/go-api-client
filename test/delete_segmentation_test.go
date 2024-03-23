// This file is auto-generated

package test

import (
	"testing"
	"time"
)

func TestDeleteSegmentation(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'deletes segmentation'
	time.Sleep(10 * time.Second)
	resp2, err = client.NewCreatePropertyBasedSegmentation("seg1", "items", "str_property").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteSegmentation("seg1").Send()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	resp, err = client.NewDeleteSegmentation("seg1").Send()
	checkErrorResponse(t, err, 404)
}
