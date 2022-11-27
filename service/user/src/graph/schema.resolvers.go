package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"user/graph/generated"
	"user/graph/model"
	"user/sql"

	"github.com/jackc/pgx/v4"
	"github.com/nicelogic/authutil"
)

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, changes map[string]interface{}) (*model.User, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v update: %v\n", user.Id, changes)
	changesJson, err := json.Marshal(changes)
	if err != nil {
		log.Printf("json marshal changes err: %v\n", err)
		return nil, err
	}
	changesJsonString := string(changesJson)
	log.Printf("changes: %s\n", changesJsonString)
	row := r.CrdbClient.Pool.QueryRow(ctx, sql.UpsertUser, user.Id, changesJsonString)
	var id, data, name string
	var update_time time.Time
	err = row.Scan(&id, &data, &name, &update_time)
	if err != nil {
		log.Printf("scan err: %v", err)
		return nil, err
	}
	log.Printf("update_time: %v\n", update_time)
	updatedUser := &model.User{
		ID:   id,
		Name: &name,
		Data: &data,
	}
	return updatedUser, err
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v query own info\n", user.Id)
	row := r.CrdbClient.Pool.QueryRow(ctx, sql.QueryMe, user.Id)
	var id, data, name string
	var update_time time.Time
	err = row.Scan(&id, &data, &name, &update_time)
	if err == pgx.ErrNoRows {
		log.Printf("user never update his data")
	} else if err != nil {
		log.Printf("scan err: %v", err)
		return nil, err
	}
	log.Printf("update_time: %v\n", update_time)
	me := &model.User{
		ID:   user.Id,
		Name: &name,
		Data: &data,
	}
	return me, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, idOrName string) ([]*model.User, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v query idOrName: %s\n", user.Id, idOrName)
	userSlice, err := r.CrdbClient.Query(ctx, sql.QueryNameOrId, idOrName)
	if err != nil {
		log.Printf("query users(idOrName: %s) err: %v\n", idOrName, err)
		return nil, err
	}
	queriedUsers := make([]*model.User, 0)
	for _, plainUser := range userSlice {
		plainUser := plainUser.([]any)
		queriedUser := &model.User{}
		queriedUser.ID = plainUser[0].(string)
		data := plainUser[1].(map[string]interface{})
		log.Printf("qeuried user (%v), data(%v)\n", queriedUser.ID, data)
		dataJson, _ := json.Marshal(data)
		dataJsonString := string(dataJson)
		queriedUser.Data = &dataJsonString
		name := plainUser[2].(string)
		queriedUser.Name = &name
		updateTime := plainUser[3].(time.Time)
		log.Printf("qeuried user(%s), data(%s), name(%s), updateTime: %v\n",
			queriedUser.ID,
			*queriedUser.Data,
			*queriedUser.Name,
			updateTime)
		queriedUsers = append(queriedUsers, queriedUser)
	}
	return queriedUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
