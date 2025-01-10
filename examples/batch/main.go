package main

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"github.com/recombee/go-api-client/v5/recombee/requests"
	"os"
)

func main() {
	client, err := recombee.NewRecombeeClient(os.Getenv("DB_ID"), os.Getenv("PRIVATE_TOKEN"), "eu-west")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Batch can encapsulate requests of various types
	batchReq := client.NewBatch([]requests.Request{
		client.NewAddDetailView("user-x", "item-y").SetCascadeCreate(true),
		client.NewRecommendItemsToUser("user-x", 5).SetScenario("just_for_you").SetCascadeCreate(true),
		client.NewRecommendItemsToItem("item-y", "user-x", 5).SetCascadeCreate(true),
		client.NewSetItemValues("item-y", map[string]interface{}{"price": 200, "category": "furniture"}),
	})

	batchResponse, err := batchReq.SendWithContext(context.Background())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Check all requests succeeded
	for i, resp := range batchResponse {
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			fmt.Printf("Request #%d failed with status code %d: %s\n", i, resp.StatusCode, resp.Error.ErrorMessage)
		}
	}

	// Print recommended items
	for _, rec := range batchResponse[1].Result.(*bindings.RecommendationResponse).Recomms {
		fmt.Println(rec.Id)
	}
}
