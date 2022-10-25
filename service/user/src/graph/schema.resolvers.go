package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"user/cassandra"
	"user/graph/generated"
	"user/graph/model"

	"github.com/nicelogic/auth"
	"golang.org/x/exp/maps"
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) UpdateUser(ctx context.Context, changes map[string]interface{}) (responseUser *model.User, err error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return
	}
	fmt.Printf("user: %v update: %v\n", user.Id, changes)

	gql, variables, err := cassandra.UpdateUserGql(changes)
	if err != nil {
		return
	}
	variables["id"] = user.Id
	response, err := r.CassandraClient.Mutation(gql, variables)
	if err != nil {
		return
	}
	fmt.Printf("response: %v\n", response)
	responseUser, err = cassandra.GetUpdatedUserFromResponse(response)
	if err != nil {
		return
	}
	return
}

func (r *queryResolver) Me(ctx context.Context) (me *model.User, err error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return
	}
	fmt.Printf("user: %v query own info\n", user.Id)

	variables := map[string]interface{}{
		"id": user.Id,
	}
	response, err := r.CassandraClient.Query(cassandra.QueryUserByIdGql, variables)
	if err != nil {
		return
	}
	fmt.Printf("response: %v\n", response)
	me, err = cassandra.GetUserFromQueryUserByIdResponse(response)
	if err != nil {
		return
	} else if me == nil {
		err = errors.New("not found me")
		return
	}
	return
}

func (r *queryResolver) Users(ctx context.Context, idOrName string) (users []*model.User, err error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return
	}
	fmt.Printf("user: %v query idOrName: %s\n", user.Id, idOrName)

	variables := map[string]interface{}{
		"id": idOrName,
	}
	response, err := r.CassandraClient.Query(cassandra.QueryUserByIdGql, variables)
	if err != nil {
		return
	}
	fmt.Printf("query user by id response: %v\n", response)

	mapUsers := make(map[string]*model.User)
	queriedByIdUser, err := cassandra.GetUserFromQueryUserByIdResponse(response)
	if err != nil {
		return
	} else if queriedByIdUser != nil {
		mapUsers[queriedByIdUser.ID] = queriedByIdUser
		fmt.Printf("queried user by id: %v\n", queriedByIdUser)
	}

	variables = map[string]interface{}{
		"name": idOrName,
	}
	for {
		response, err = r.CassandraClient.Query(cassandra.QueryUserByNameGql, variables)
		if err != nil {
			return
		}
		fmt.Printf("query user by name response: %v\n", response)

		queriedUsers, pageState, getErr := cassandra.GetUserFromQueryUserByNameResponse(response)
		if getErr != nil {
			err = getErr
			return
		}
		for id, user := range queriedUsers {
			mapUsers[id] = user
		}
		if pageState == "" {
			break
		} else {
			variables["pageState"] = pageState
		}
	}

	users = maps.Values(mapUsers)
	return
}
