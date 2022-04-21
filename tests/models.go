package main

import "github.com/hasura/go-graphql-client"

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
