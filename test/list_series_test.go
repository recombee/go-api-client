package test

import (
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"testing"
	"time"
)

func TestListSeries(t *testing.T) {
	var resp []bindings.Series
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'lists entities'
	time.Sleep(10 * time.Second)
	resp, err = client.NewListSeries().Send()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "entity_id", resp[0].SeriesId)
}
