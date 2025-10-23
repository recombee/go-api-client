// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
)

func TestRecommendItemSegmentsToItemSegment(t *testing.T) {
	var resp bindings.RecommendationResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	client := createClient(t)

	// it 'rejects request to scenario which is not set up'
	resp, err = client.NewRecommendItemSegmentsToItemSegment("segment_id", "entity_id", 5).SetScenario("scenario1").SetCascadeCreate(true).Send()
	checkErrorResponse(t, err, 400)
}
