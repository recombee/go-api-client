// This file is auto-generated

package test

import (
	"context"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"testing"
)

func TestRecommendUsersToItemInBatch(t *testing.T) {
	defaultTestSetUp(t)
	recommsTestSetUp(t)
	// it 'recommends'
	// it 'recommends to previously nonexisting entity with cascadeCreate'
	// it 'recommends with expert settings'
	// it 'recommends with reql expressions'
	client := createClient(t)
	var resp2 interface{}
	var err error
	useVars(resp2, err)

	var reqs = []requests.Request{
		client.NewRecommendUsersToItem("entity_id", 9),
		client.NewRecommendUsersToItem("nonexisting", 9).SetCascadeCreate(true),
		client.NewRecommendUsersToItem("nonexisting2", 9).SetCascadeCreate(true).SetExpertSettings(map[string]interface{}{}),
		client.NewRecommendUsersToItem("nonexisting2", 9).SetCascadeCreate(true).SetReqlExpressions(map[string]string{"boolean": "true", "number": "if ('answer' > 0) then 1 else 2", "string": "\"test\""}),
	}

	batchResponse, err := client.NewBatch(reqs).SendWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, 200, batchResponse[0].StatusCode)
	assertEqual(t, 200, batchResponse[1].StatusCode)
	assertEqual(t, 200, batchResponse[2].StatusCode)
	assertEqual(t, 200, batchResponse[3].StatusCode)
}
