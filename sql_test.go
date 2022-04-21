package main

import (
	"context"
	"fmt"
	"testing"

	graphql "github.com/hasura/go-graphql-client"
	"github.com/sammychinedu2ky/synthesis/graph"
)

func TestCalculate(t *testing.T) {
	if graph.Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculate1(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		playersFromId struct {
			playerid graphql.ID
			username graphql.String
		} `graphql:"playersFromId(playerid: $playerid)"`
	}
	variables := map[string]interface{}{
		"playerid": graphql.ID("1"),
	}
	err := client.Query(context.Background(), &q, variables)

	fmt.Println(q)
	if err != nil {
		fmt.Println(err)
		return
	}
	if graph.Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
