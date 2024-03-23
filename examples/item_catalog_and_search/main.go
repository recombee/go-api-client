package main

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/recombee"
	"github.com/recombee/go-api-client/recombee/requests"
	"math/rand"
	"os"
	"time"
)

func main() {
	client, err := recombee.NewRecombeeClient(os.Getenv("DB_ID"), os.Getenv("PRIVATE_TOKEN"), "eu-west")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Add item properties
	// Properties are the columns of the items catalogs
	properties := []struct {
		propertyName string
		propertyType string
	}{
		{"price", "double"},
		{"num-cores", "int"},
		{"description", "string"},
		{"image", "image"},
	}

	for _, p := range properties {
		if _, err := client.NewAddItemProperty(p.propertyName, p.propertyType).Send(); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	// Prepare and send catalog of computers
	const numItems = 100
	requestsBatch := make([]requests.Request, numItems)

	for i := 0; i < numItems; i++ {
		itemId := fmt.Sprintf("computer-%d", i)
		requestsBatch[i] = client.NewSetItemValues(itemId, map[string]interface{}{
			"price":       600.0 + 400*rand.Float64(),
			"num-cores":   1 + rand.Intn(7),
			"description": "Great computer",
			"image":       fmt.Sprintf("http://examplesite.com/products/%s.jpg", itemId),
		}).SetCascadeCreate(true)
	}

	if _, err := client.NewBatch(requestsBatch).Send(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Simulate random purchases
	const probabilityPurchased = 0.02
	var addPurchaseRequests []requests.Request

	for i := 0; i < numItems; i++ {
		for j := 0; j < numItems; j++ {
			if rand.Float64() < probabilityPurchased {
				userId := fmt.Sprintf("user-%d", i)
				itemId := fmt.Sprintf("computer-%d", j)
				addPurchaseRequests = append(addPurchaseRequests, client.NewAddPurchase(userId, itemId).SetCascadeCreate(true))
			}
		}
	}

	// Send the Batch request with custom timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := client.NewBatch(addPurchaseRequests).SendWithContext(ctx); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Request recommendations with a set scenario
	// Behavior of the recommendations is configured per scenario in the Recombee Admin UI
	req := client.NewRecommendItemsToItem("computer-6", "user-42", 5).SetScenario("product_detail")

	recommRes, err := req.Send()
	if err != nil {
		panic(err)
	}
	fmt.Println("Recommended items:")
	for _, rec := range recommRes.Recomms {
		fmt.Println(rec.Id)
	}

	// Recommend items with at least 3 processor cores (dynamic filter), return property values of the items
	req = client.NewRecommendItemsToItem("computer-6", "user-42", 5)
	req = req.SetFilter("'num-cores'>=3").SetReturnProperties(true)
	recommRes, err = req.Send()
	fmt.Println("\nRecommended items with at least 3 cores:")
	for _, rec := range recommRes.Recomms {
		fmt.Println(rec.Id)
		fmt.Println(rec.Values)
	}

	// Perform a search using "computers" query
	searchResponse, err := client.NewSearchItems("user-42", "computers", 5).Send()

	if err != nil {
		panic(err)
	}
	fmt.Println("\nSearch matches: %s", searchResponse.Recomms)
}
