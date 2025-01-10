// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"testing"
)

func TestSearchItemSegments(t *testing.T) {
	var resp bindings.SearchResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'rejects request to scenario which is not set up'
	resp, err = client.NewSearchItemSegments("entity_id", "query", 5).SetScenario("scenario1").SetCascadeCreate(true).Send()
	checkErrorResponse(t, err, 400)
}
