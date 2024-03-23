// This file is auto-generated

package test

import (
	"testing"
)

func TestUpdateAutoReqlSegmentation(t *testing.T) {
	var resp string
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'updates auto ReQL segmentation'
	resp2, err = client.NewCreateAutoReqlSegmentation("seg1", "items", "{'str_property'}").SetTitle("Test Segmentation").SetDescription("For test purposes").Send()
	if err != nil {
		t.Fatal(err)
	}
	resp, err = client.NewUpdateAutoReqlSegmentation("seg1").SetTitle("New title").SetExpression("{'str_property' + 'str_property'}").SetDescription("Updated").Send()
	if err != nil {
		t.Fatal(err)
	}
}
