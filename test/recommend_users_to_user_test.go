// This file is auto-generated

package test

import (
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"testing"
)

func TestRecommendUsersToUser(t *testing.T) {
	var resp bindings.RecommendationResponse
	var resp2 interface{}
	var err error
	useVars(resp, resp2, err)

	defaultTestSetUp(t)
	recommsTestSetUp(t)
	client := createClient(t)

	// it 'recommends'
	resp, err = client.NewRecommendUsersToUser("entity_id", 9).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'recommends to previously nonexisting entity with cascadeCreate'
	resp, err = client.NewRecommendUsersToUser("nonexisting", 9).SetCascadeCreate(true).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'recommends with expert settings'
	resp, err = client.NewRecommendUsersToUser("nonexisting2", 9).SetCascadeCreate(true).SetExpertSettings(map[string]interface{}{}).Send()
	if err != nil {
		t.Fatal(err)
	}

	// it 'recommends with reql expressions'
	resp, err = client.NewRecommendUsersToUser("nonexisting2", 9).SetCascadeCreate(true).SetReqlExpressions(map[string]string{"boolean": "true", "number": "if ('answer' > 0) then 1 else 2", "string": "\"test\""}).Send()
	if err != nil {
		t.Fatal(err)
	}
}
