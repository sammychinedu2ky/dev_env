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

type Cohorts struct {
	Cohortid          graphql.ID
	Name              graphql.String
	Date              graphql.String
	At_sku            graphql.Int
	At_id             graphql.String
	Status            graphql.String
	Zoom              graphql.String
	Game              graphql.String
	Level             graphql.Int
	Imported_students graphql.Int
	Start_date        graphql.String
	End_date          graphql.String
	Mapaccess         graphql.Int
	Record            graphql.Int
	Event_id          graphql.String
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
	player1 := Players{
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
	if q.PlayersFromId == player1 {

	} else {
		t.Error("Error! Expected value is ", player1)
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
	player1 := Players{
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
		t.Error("Error! Expected value is ", player1)
		return
	}
	if q.PlayersFromUserName == player1 {

	} else {
		t.Error("Error! Expected value is ", player1)
	}
}

func TestPlayers(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		Players []struct {
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
		} `graphql:"players"`
	}
	err := client.Query(context.Background(), &q, nil)
	player1 := Players{
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
	player2 := Players{
		"2",
		"6",
		"Mark Burton",
		"congue@google.org",
		16,
		"",
		"jifo_41",
		"EWQ44QUN5VT",
		1,
		216,
		"HYP64VNY2ON",
		"2021-05-14T19:34:06Z",
		"2022-09-17T02:41:21Z",
		"Colombia/ChocÃ³",
		"abc",
		"25266604794901524264AB9387",
		"2022-02-02T15:15:04Z",
		"3",
	}

	fmt.Println(q.Players[0])
	if err != nil {
		fmt.Println(err)
		t.Error("Error! Expected value is ", player1, player2)
		return
	}
	if q.Players[0] == player1 && q.Players[1] == player2 {
	} else {
		t.Error("Error! Expected value is ", player1, player2)
	}
}

func TestCohorts(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		Cohorts []struct {
			Cohortid          graphql.ID
			Name              graphql.String
			Date              graphql.String
			At_sku            graphql.Int
			At_id             graphql.String
			Status            graphql.String
			Zoom              graphql.String
			Game              graphql.String
			Level             graphql.Int
			Imported_students graphql.Int
			Start_date        graphql.String
			End_date          graphql.String
			Mapaccess         graphql.Int
			Record            graphql.Int
			Event_id          graphql.String
		} `graphql:"cohorts"`
	}
	err := client.Query(context.Background(), &q, nil)
	cohort1 := Cohorts{
		"1",
		"Brynne",
		"",
		13,
		"abp",
		"incomplete",
		"no",
		"porttitor",
		8,
		1,
		"2021-04-21T13:03:30Z",
		"2021-11-30T04:11:51Z",
		1,
		10,
		"IDX_307e69e06a2315f03bdc9b3fec",
	}
	cohort2 := Cohorts{
		"2",
		"Sasha",
		"",
		11,
		"abs",
		"complete",
		"yes",
		"mattis.",
		6,
		2,
		"2021-06-19T15:05:38Z",
		"2021-12-06T11:03:36Z",
		0,
		15,
		"IDX_302e69e06a8055f03bdc9b3fec",
	}
	if err != nil {
		fmt.Println(err)
		t.Error("Error! Expected value is ", cohort1, cohort2)
		return
	}
	if q.Cohorts[0] == cohort1 && q.Cohorts[1] == cohort2 {
	} else {
		t.Error("Error! Expected value is ", cohort1, cohort2)
	}
}
