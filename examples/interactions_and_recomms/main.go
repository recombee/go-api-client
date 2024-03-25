package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/recombee/go-api-client/v4/recombee"
	"github.com/recombee/go-api-client/v4/recombee/requests"
)

func main() {
	client, err := recombee.NewRecombeeClient(os.Getenv("DB_ID"), os.Getenv("PRIVATE_TOKEN"), "eu-west")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	const numItemsAndUsers = 100
	const probabilityInteracted = 0.1

	var addDetailViewInteractions []requests.Request
	for i := 0; i < numItemsAndUsers; i++ {
		for j := 0; j < numItemsAndUsers; j++ {
			if rand.Float64() < probabilityInteracted {
				userId := fmt.Sprintf("user-%d", i)
				itemId := fmt.Sprintf("item-%d", j)

				request := client.NewAddDetailView(userId, itemId).SetCascadeCreate(true)
				addDetailViewInteractions = append(addDetailViewInteractions, request)
			}
		}
	}

	fmt.Println("Send all interactions in Batch")
	batchRes, err := client.NewBatch(addDetailViewInteractions).Send()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(batchRes) // Contains responses for all requests

	// Get 5 recommendations for user 'user-25'
	recommendReq := client.NewRecommendItemsToUser("user-25", 5).SetCascadeCreate(true)
	recommendRes, err := recommendReq.Send()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Recommended items:")
	for _, rec := range recommendRes.Recomms {
		fmt.Println(rec.Id)
	}

	// User scrolled down - get next 3 recommended items
	fmt.Println("Next recommended items:")
	recommendNextRes, err := client.NewRecommendNextItems(recommendRes.RecommId, 3).Send()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, rec := range recommendNextRes.Recomms {
		fmt.Println(rec.Id)
	}
}
