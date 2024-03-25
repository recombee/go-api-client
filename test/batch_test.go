package test

import (
	"fmt"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"github.com/recombee/go-api-client/v4/recombee/requests"
	"testing"
)

func TestBatch(t *testing.T) {
	defaultTestSetUp(t)
	client := createClientWithOptions(t)
	const numItemsAndUsers = 100

	var reqs []requests.Request
	for i := 0; i < numItemsAndUsers; i++ {
		for j := 0; j < numItemsAndUsers; j++ {
			userId := fmt.Sprintf("user-%d", i)
			itemId := fmt.Sprintf("item-%d", j)

			request := client.NewAddDetailView(userId, itemId).SetCascadeCreate(true)
			reqs = append(reqs, request)
		}
	}
	reqs = append(reqs, client.NewRecommendItemsToUser("tst_user", 9).SetCascadeCreate(true))

	batchReq := client.NewBatch(reqs)
	batchReq.MaxRequestsPerBatchCall = 50 // Set small chunk size to test multiple batch calls

	batchResponse, err := batchReq.Send()
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, len(batchResponse), len(reqs))

	for _, resp := range batchResponse {
		assertEqual(t, 200, resp.StatusCode)
	}

	// Check the response of the last request
	recommendationResponse := batchResponse[len(batchResponse)-1].Result.(*bindings.RecommendationResponse)
	assertEqual(t, 9, len(recommendationResponse.Recomms))
}
