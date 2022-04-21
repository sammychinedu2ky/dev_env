package main

import (
	"context"
	"fmt"
	"testing"

	graphql "github.com/hasura/go-graphql-client"
)

func TestCalculate1(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		PlayersFromId struct {
			Playerid graphql.ID
			Username graphql.String
		} `graphql:"playersFromId(playerid: \"1\")"`
	}
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	if q.PlayersFromId.Playerid == "1" && q.PlayersFromId.Username == "lwgj_45" {
	} else {
		t.Error("Error! Expected {{1,lwgj_45}}")
	}
}
