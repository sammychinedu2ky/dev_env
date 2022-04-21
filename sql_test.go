package main

import (
	"context"
	"fmt"
	"testing"

	graphql "github.com/hasura/go-graphql-client"
)

type Players struct {
	Playerid        graphql.ID
	Cohortid        graphql.ID
	Name            graphql.String
	Email           graphql.String
	Age             graphql.Int
	Dob             graphql.String
	Username        graphql.String
	Password        graphql.String
	Admin           graphql.Int
	Order           graphql.Int
	Reset_token     graphql.String
	Password_sent   graphql.String
	Last_login      graphql.String
	Timezone        graphql.String
	At_id           graphql.String
	Hashed_password graphql.String
	Start_date      graphql.String
	User_id         graphql.String
}

func TestPlayersFromID(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		PlayersFromId struct {
			Playerid        graphql.ID
			Cohortid        graphql.ID
			Name            graphql.String
			Email           graphql.String
			Age             graphql.Int
			Dob             graphql.String
			Username        graphql.String
			Password        graphql.String
			Admin           graphql.Int
			Order           graphql.Int
			Reset_token     graphql.String
			Password_sent   graphql.String
			Last_login      graphql.String
			Timezone        graphql.String
			At_id           graphql.String
			Hashed_password graphql.String
			Start_date      graphql.String
			User_id         graphql.String
		} `graphql:"playersFromId(playerid: \"1\")"`
	}
	err := client.Query(context.Background(), &q, nil)
	ans := Players{
		"1",
		"3",
		"Leroy Myers",
		"orci@outlook.edu",
		25,
		"",
		"lwgj_45",
		"RMK21YVP2BT",
		1,
		192,
		"CWZ23SXQ2LJ",
		"2022-05-16T19:21:50Z",
		"2022-01-24T04:31:20Z",
		"Chile/Los RÃ­os",
		"abs",
		"10734068542471624377AB8448",
		"2022-02-19T21:20:02Z",
		"1",
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	if q.PlayersFromId == ans {

	} else {
		t.Error("Error! Expected value is ", ans)
	}
}

func TestPlayersFromUserName(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		PlayersFromUserName struct {
			Playerid        graphql.ID
			Cohortid        graphql.ID
			Name            graphql.String
			Email           graphql.String
			Age             graphql.Int
			Dob             graphql.String
			Username        graphql.String
			Password        graphql.String
			Admin           graphql.Int
			Order           graphql.Int
			Reset_token     graphql.String
			Password_sent   graphql.String
			Last_login      graphql.String
			Timezone        graphql.String
			At_id           graphql.String
			Hashed_password graphql.String
			Start_date      graphql.String
			User_id         graphql.String
		} `graphql:"playersFromId(playerid: \"lwgj_45\")"`
	}
	err := client.Query(context.Background(), &q, nil)
	ans := Players{
		"1",
		"3",
		"Leroy Myers",
		"orci@outlook.edu",
		25,
		"",
		"lwgj_45",
		"RMK21YVP2BT",
		1,
		192,
		"CWZ23SXQ2LJ",
		"2022-05-16T19:21:50Z",
		"2022-01-24T04:31:20Z",
		"Chile/Los RÃ­os",
		"abs",
		"10734068542471624377AB8448",
		"2022-02-19T21:20:02Z",
		"1",
	}

	if err != nil {
		fmt.Println(err)
		t.Error("Error! Expected value is ", ans)
		return
	}
	if q.PlayersFromUserName == ans {

	} else {
		t.Error("Error! Expected value is ", ans)
	}
}
