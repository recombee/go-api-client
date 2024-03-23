package test

import (
	"fmt"
	"github.com/recombee/go-api-client/recombee"
	"github.com/recombee/go-api-client/recombee/errors"
	"github.com/recombee/go-api-client/recombee/requests"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

func createClient(t *testing.T) *recombee.RecombeeClient {
	client, err := recombee.NewRecombeeClient(os.Getenv("DB_ID"), os.Getenv("PRIVATE_TOKEN"), "eu-west")
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}
	return client
}

func createClientWithOptions(t *testing.T) *recombee.RecombeeClient {
	apiUri := "rapi-eu-west.recombee.com"
	options := recombee.ClientOptions{BaseUri: &apiUri, HttpClient: &http.Client{}}
	client, err := recombee.NewRecombeeClientWithOptions(os.Getenv("DB_ID"), os.Getenv("PRIVATE_TOKEN"), options)
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}
	return client
}

func checkErrorResponse(t *testing.T, err error, expectedStatusCode int) {
	if err == nil {
		t.Fatal("Expected an error, got none")
	}

	respErr, ok := err.(*errors.ResponseError)
	if !ok {
		t.Fatalf("Expected a ResponseError, got %T", err)
	}
	if respErr.StatusCode != expectedStatusCode {
		t.Fatalf("Expected status code %d, got %d", expectedStatusCode, respErr.StatusCode)
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected == nil || actual == nil {
		if reflect.DeepEqual(expected, actual) {
			return
		}
		t.Fatalf("Expected and actual values are not equal.\nExpected: %v\nActual:   %v", expected, actual)
		return
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	if expectedValue.Kind() != actualValue.Kind() {
		if bothAreNumbers(expectedValue, actualValue) {
			if numbersAreEqual(expectedValue, actualValue) {
				return // Numbers are considered equal
			}
		}
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected and actual values are not equal.\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func bothAreNumbers(a, b reflect.Value) bool {
	return isNumber(a) && isNumber(b)
}

func isNumber(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// numbersAreEqual converts both numbers to float64 and compares them.
func numbersAreEqual(a, b reflect.Value) bool {
	return a.Convert(reflect.TypeOf(float64(0))).Float() == b.Convert(reflect.TypeOf(float64(0))).Float()
}

func useVars(vars ...interface{}) {}

func defaultTestSetUp(t *testing.T) {
	client := createClient(t)

	client.NewResetDatabase().Send()

	for {
		_, err := client.NewListItems().Send()
		if err == nil {
			break
		}
		if errResp, ok := err.(*errors.ResponseError); ok && errResp.StatusCode != 422 {
			t.Fatalf("Unexpected error: %v", err)
		}
	}

	batchReq := client.NewBatch([]requests.Request{
		client.NewAddItem("entity_id"),
		client.NewAddUser("entity_id"),
		client.NewAddSeries("entity_id"),
		client.NewInsertToSeries("entity_id", "item", "entity_id", 1),
		client.NewAddItemProperty("int_property", "int"),
		client.NewAddItemProperty("str_property", "string"),
		client.NewSetItemValues("entity_id", map[string]interface{}{"int_property": 42, "str_property": "hello"}),
		client.NewAddUserProperty("int_property", "int"),
		client.NewAddUserProperty("str_property", "string"),
		client.NewSetUserValues("entity_id", map[string]interface{}{"int_property": 42, "str_property": "hello"}),
	})

	batchReq.Send()
}

func interactionsTestSetUp(t *testing.T) {

	client := createClient(t)

	batchReq := client.NewBatch([]requests.Request{
		client.NewAddUser("user"),
		client.NewAddItem("item"),
		client.NewAddDetailView("user", "item"),
		client.NewAddPurchase("user", "item").SetTimestamp(time.Unix(0, 0)),
		client.NewAddRating("user", "item", 1).SetTimestamp(time.Unix(0, 0)),
		client.NewAddCartAddition("user", "item").SetTimestamp(time.Unix(0, 0)),
		client.NewAddBookmark("user", "item").SetTimestamp(time.Unix(0, 0)),
		client.NewSetViewPortion("user", "item", 1).SetTimestamp(time.Unix(0, 0)),
	})
	batchReq.Send()
}

func recommsTestSetUp(t *testing.T) {
	client := createClient(t)

	num := 1000
	probabilityPurchased := 0.007

	rqs := []requests.Request{client.NewAddItemProperty("answer", "int"),
		client.NewAddItemProperty("id2", "string"),
		client.NewAddItemProperty("empty", "string"),
	}

	myUsers := make([]string, num)
	myItems := make([]string, num)

	for i := 0; i < num; i++ {
		myUsers[i] = fmt.Sprintf("user-%d", i)
		myItems[i] = fmt.Sprintf("item-%d", i)

		rqs = append(rqs, client.NewAddUser(myUsers[i]))
		rqs = append(rqs, client.NewSetItemValues(myItems[i], map[string]interface{}{"answer": 42, "id2": myItems[i]}).SetCascadeCreate(true))
	}

	for _, user := range myUsers {
		for _, item := range myItems {
			if rand.Float64() < probabilityPurchased {
				rqs = append(rqs, client.NewAddPurchase(user, item))
			}
		}
	}

	client.NewBatch(rqs).Send()
}
