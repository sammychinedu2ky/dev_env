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

type Simulations struct {
	Simulationid  graphql.ID
	Name          graphql.String
	Game_url      graphql.String
	Setup_url     graphql.String
	Env           graphql.String
	Configvar     graphql.String
	Field_mapping graphql.String
}

type StarSimulations struct {
	Id           graphql.ID
	Name         graphql.String
	GameUrl      graphql.String
	SetupUrl     graphql.String
	Env          graphql.String
	Configvar    graphql.String
	FieldMapping graphql.String
}

type GameRepositories struct {
	Gameid   graphql.ID
	Token    graphql.String
	Color    graphql.String
	Gametype graphql.String
	Teamid   graphql.ID
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
		} `graphql:"playersFromUserName(username: \"lwgj_45\")"`
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

func TestCohortsFromId(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		CohortsFromId struct {
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
		} `graphql:"cohortsFromId(cohortid: \"1\")"`
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
	if err != nil {
		fmt.Println(err)
		t.Error("Error! Expected value is ", cohort1)
		return
	}
	if q.CohortsFromId == cohort1 {
	} else {
		t.Error("Error! Expected value is ", cohort1)
	}
}

func TestSimulations(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		Simulations []struct {
			Simulationid  graphql.ID
			Name          graphql.String
			Game_url      graphql.String
			Setup_url     graphql.String
			Env           graphql.String
			Configvar     graphql.String
			Field_mapping graphql.String
		} `graphql:"simulations"`
	}
	err := client.Query(context.Background(), &q, nil)
	simulation1 := Simulations{
		"1",
		"Gavin",
		"https://OZS51GPN2RF.in",
		"https://LER11DIC2LO.eu",
		"77SKWQ24EDX3EG",
		"Q8OEJT371G9",
		"sczixcnh",
	}
	simulation2 := Simulations{
		"2",
		"Troy",
		"https://UWP39GIQ8GY.in",
		"https://MDQ06HPL7ZB.eu",
		"16NVBT54YBB1PN",
		"T5MLOE711P7",
		"ftvcgdne",
	}
	if err != nil {
		fmt.Println(err)
		t.Error("Error! Expected value is ", simulation1, simulation2)
		return
	}
	if q.Simulations[0] == simulation1 && q.Simulations[1] == simulation2 {
	} else {
		t.Error("Error! Expected value is ", simulation1, simulation2)
	}
}

func TestStarsSimulations(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		StarsSimulation struct {
			Id           graphql.ID
			Name         graphql.String
			GameUrl      graphql.String
			SetupUrl     graphql.String
			Env          graphql.String
			Configvar    graphql.String
			FieldMapping graphql.String
		} `graphql:"starsSimulation(playerid: \"1\")"`
	}
	err := client.Query(context.Background(), &q, nil)
	starSimulation1 := StarSimulations{
		"5",
		"Melvin",
		"https://FSC57LJD8LF.in",
		"https://DHU88PMH5JN.eu",
		"82OIHY45BWU1IL",
		"B9DVCX934M3",
		"oywvgmuf",
	}
	if err != nil {
		fmt.Println(err)
		t.Error("Error! Expected value is ", starSimulation1)
		return
	}
	if q.StarsSimulation == starSimulation1 {
	} else {
		t.Error("Error! Expected value is ", starSimulation1)
	}
}

func TestGameRepository(t *testing.T) {
	client := graphql.NewClient("http://localhost:3000/query", nil)
	var q struct {
		GameRepository struct {
			Gameid   graphql.ID
			Token    graphql.String
			Color    graphql.String
			Gametype graphql.String
			Teamid   graphql.ID
		} `graphql:"gameRepository(playerid: \"1\",cohortid: \"1\")"`
	}
	err := client.Query(context.Background(), &q, nil)
	gameRepository1 := GameRepositories{
		"",
		"",
		"",
		"",
		"",
	}
	if err != nil {
	}
	if q.GameRepository == gameRepository1 {
	} else {
		t.Error("Error! Expected value is ", gameRepository1)
	}
}
