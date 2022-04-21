package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sammychinedu2ky/synthesis/graph/generated"
	"github.com/sammychinedu2ky/synthesis/graph/model"
	"github.com/vektah/gqlparser/gqlerror"
)

func (r *queryResolver) PlayersFromID(ctx context.Context, playerid string) (*model.Players, error) {
	//panic(fmt.Errorf("not implemented"))
	var player model.Players
	query := "SELECT admin,age,at_id,cohortid,dob,email,hashed_password,last_login,name,players.order,password,password_sent,playerid,reset_token,start_date,timezone,user_id,username FROM players where playerid = ?"

	err := r.Db.QueryRow(query, playerid).Scan(&player.Admin, &player.Age, &player.AtID, &player.Cohortid, &player.Dob,
		&player.Email, &player.HashedPassword, &player.LastLogin, &player.Name, &player.Order,
		&player.Password, &player.PasswordSent, &player.Playerid, &player.ResetToken,
		&player.StartDate, &player.Timezone, &player.UserID, &player.Username)
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}

	return &player, nil
}

func (r *queryResolver) PlayersFromUserName(ctx context.Context, username string) (*model.Players, error) {
	//panic(fmt.Errorf("not implemented"))
	var player model.Players
	query := "SELECT admin,age,at_id,cohortid,dob,email,hashed_password,last_login,name,players.order,password,password_sent,playerid,reset_token,start_date,timezone,user_id,username FROM players where playerid = ?"

	err := r.Db.QueryRow(query, username).Scan(&player.Admin, &player.Age, &player.AtID, &player.Cohortid, &player.Dob,
		&player.Email, &player.HashedPassword, &player.LastLogin, &player.Name, &player.Order,
		&player.Password, &player.PasswordSent, &player.Playerid, &player.ResetToken,
		&player.StartDate, &player.Timezone, &player.UserID, &player.Username)
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}

	return &player, nil
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Players, error) {
	//panic(fmt.Errorf("not implemented"))
	players := []*model.Players{}
	query := "SELECT admin,age,at_id,cohortid,dob,email,hashed_password,last_login,name,players.order,password,password_sent,playerid,reset_token,start_date,timezone,user_id,username FROM players"
	res, err := r.Db.Query(query)
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}
	for res.Next() {
		player := model.Players{}
		err := res.Scan(&player.Admin, &player.Age, &player.AtID, &player.Cohortid, &player.Dob,
			&player.Email, &player.HashedPassword, &player.LastLogin, &player.Name, &player.Order,
			&player.Password, &player.PasswordSent, &player.Playerid, &player.ResetToken,
			&player.StartDate, &player.Timezone, &player.UserID, &player.Username)

		if err != nil {
			fmt.Println(err)
			graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
		}
		players = append(players, &player)
	}
	return players, nil
}

func (r *queryResolver) Cohorts(ctx context.Context) ([]*model.Cohort, error) {
	//panic(fmt.Errorf("not implemented"))
	cohorts := []*model.Cohort{}
	res, err := r.Db.Query("SELECT at_id,at_sku,cohortid,date,end_date,event_id,game,imported_students,level,mapaccess,name,record,start_date,status,zoom FROM cohorts")
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}
	for res.Next() {
		cohort := model.Cohort{}
		err := res.Scan(&cohort.AtID, &cohort.AtSku,
			&cohort.Cohortid, &cohort.Date, &cohort.EndDate,
			&cohort.EventID, &cohort.Game, &cohort.ImportedStudents,
			&cohort.Level, &cohort.Mapaccess, &cohort.Name, &cohort.Record,
			&cohort.StartDate, &cohort.Status, &cohort.Zoom)

		if err != nil {
			fmt.Println(err)
			graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
		}
		cohorts = append(cohorts, &cohort)
	}
	return cohorts, nil
}

func (r *queryResolver) CohortsFromID(ctx context.Context, cohortid string) (*model.Cohort, error) {
	//panic(fmt.Errorf("not implemented"))
	var cohort model.Cohort

	err := r.Db.QueryRow("SELECT at_id,at_sku,cohortid,date,end_date,event_id,game,imported_students,level,mapaccess,name,record,start_date,status,zoom FROM cohorts where cohortid = ?;", cohortid).Scan(&cohort.AtID, &cohort.AtSku,
		&cohort.Cohortid, &cohort.Date, &cohort.EndDate,
		&cohort.EventID, &cohort.Game, &cohort.ImportedStudents,
		&cohort.Level, &cohort.Mapaccess, &cohort.Name, &cohort.Record,
		&cohort.StartDate, &cohort.Status, &cohort.Zoom)

	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}

	return &cohort, nil
}

func (r *queryResolver) Simulations(ctx context.Context) ([]*model.Simulations, error) {
	simulations := []*model.Simulations{}
	res, err := r.Db.Query("SELECT configvar, env, field_mapping, game_url,name,setup_url,simulationid from simulations")
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}
	for res.Next() {
		simulation := model.Simulations{}
		err := res.Scan(&simulation.Configvar, &simulation.Env, &simulation.FieldMapping, &simulation.GameURL,
			&simulation.Name, &simulation.SetupURL, &simulation.Simulationid)

		if err != nil {
			fmt.Println(err)
			graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
		}
		simulations = append(simulations, &simulation)
	}
	return simulations, nil
}

func (r *queryResolver) GameRepository(ctx context.Context, cohortid string, playerid string) (*model.GameRepository, error) {
	//panic(fmt.Errorf("not implemented"))
	query := "select t.color,t.gameid,g.gametype,t.teamid,t.token from stars.games g join stars.teams t using(gameid) join stars.teamplayers tp using(teamid) join stars.cohorts c on c.cohortid=g.cohortid where g.cohortid=? and g.current='1' and tp.playerid=? order by g.gameid desc limit 1"

	var gameRepository model.GameRepository
	err := r.Db.QueryRow(query, cohortid, playerid).Scan(&gameRepository.Color, &gameRepository.Gameid, &gameRepository.Gametype, &gameRepository.Teamid, &gameRepository.Token)
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}

	return &gameRepository, nil
}

func (r *queryResolver) StarsSimulation(ctx context.Context, playerid string) (*model.StarsSimulation, error) {
	//panic(fmt.Errorf("not implemented"))
	query := "select s.configvar,s.env,s.field_mapping,s.game_url as gameUrl,s.simulationid as id,s.name,s.setup_url as setupurl from stars.cohortplayers cp join stars.cohort_setup_permissions csp using(cohortid) join stars.simulations s using(simulationid) where cp.playerid=? group by s.simulationid;"
	var starsSimulation model.StarsSimulation

	err := r.Db.QueryRow(query, playerid).Scan(&starsSimulation.Configvar, &starsSimulation.Env, &starsSimulation.FieldMapping, &starsSimulation.GameURL, &starsSimulation.ID, &starsSimulation.Name, &starsSimulation.SetupURL)
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	}

	return &starsSimulation, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) CohortsFromName(ctx context.Context, name string) (*model.Cohort, error) {
	panic(fmt.Errorf("not implemented"))
	// var cohort model.Cohort
	// 	err := r.Db.QueryRow("select username from players where playerid = ?", username).Scan(&players.Username)
	// 	if err != nil {
	// 		graphql.AddError(ctx, gqlerror.Errorf(err.Error()))
	// 	}

	// 	return &players, nil
}
