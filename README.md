# Recombee API Client

A Go client (SDK) for easy use of the [Recombee](https://www.recombee.com/) recommendation API.

If you don't have an account at Recombee yet, you can create a free account [here](https://www.recombee.com/).

Documentation of the API can be found at [docs.recombee.com](https://docs.recombee.com/).

## Installation

Run the following command in your Go project:
```
go get github.com/recombee/go-api-client/v4@v4.1.1
```

## Examples

### Basic example

The following example shows how to:
- Send many interactions within a single `Batch` request
- Request recommendations for a user using `RecommendItemsToUser` and `RecommendNextItems` endpoints

```go
package main

import (
    "fmt"
    "math/rand"
    "os"

    "github.com/recombee/go-api-client/v4/recombee"
    "github.com/recombee/go-api-client/v4/recombee/requests"
)

func main() {
    client, err := recombee.NewRecombeeClient("your-database-id", "db-private-token", "us-west")
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

```


See `examples/batch` for an extended example of using the `Batch` requests.

### Using property values

The second example shows how to:
- Specify the properties (columns) of the items catalog using `AddItemProperty` requests
- Upload their values for all the items using `SetItemValues` requests in a `Batch`
- Get recommendations based on an item using `RecommendItemsToItem`
- Search items using `SearchItems`


```go
package main

import (
    "context"
    "fmt"
    "github.com/recombee/go-api-client/v4/recombee"
    "github.com/recombee/go-api-client/v4/recombee/requests"
    "math/rand"
    "os"
    "time"
)

func main() {
    client, err := recombee.NewRecombeeClient("your-database-id", "db-private-token", "us-west")
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

    // Prepare and send a catalog of computers
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

```

## Sending Requests

A request can be sent to the Recombee API using the `Send()` method.

If you want to provide your own `context.Context` (e.g. to set your custom timeout for the request) call the `SendWithContext(ctx context.Context)` method.

## RecombeeClient Initialization

The RecombeeClient can be created using the `NewRecombeeClient(...)` function, as per previous examples.

If you want to provide your own `http.Client` use `NewRecombeeClientWithOptions`:

```go
    region := "ap-se"
    options := recombee.ClientOptions{Region: &region, HttpClient: &myHttpClient}
    client, err := recombee.NewRecombeeClientWithOptions("your-database-id", "db-private-token", options)
```

## Custom Errors

Various errors can occur while processing the request, for example, because of adding an already existing item or submitting interaction of a nonexistent user without *cascadeCreate* set to true. These errors lead to returning the `errors.ResponseError` by the `Send` method.

Another reason for producing an error is a timeout, yielding `errors.TimeoutError`.

We are doing our best to provide the fastest and most reliable service, but production-level applications shall implement a fallback solution since errors can always happen. The fallback might be, for example, showing the most popular items from the current category, or not displaying recommendations at all.
