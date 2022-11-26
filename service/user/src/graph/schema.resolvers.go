package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"user/graph/generated"
	"user/graph/model"

	"github.com/nicelogic/auth"
	"github.com/nicelogic/authutil"
)

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, changes map[string]interface{}) (*model.User, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v update: %v\n", user.Id, changes)
	return nil, nil

	// gql, variables, err := cassandra.UpdateUserGql(changes)
	// if err != nil {
	// 	return nil, err
	// }
	// variables["id"] = user.Id
	// response, err := r.CassandraClient.Mutation(gql, variables)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("response: %v\n", response)
	// responseUser, err := cassandra.GetUpdatedUserFromResponse(response)
	// if err != nil {
	// 	return nil, err
	// }
	// return responseUser, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v query own info\n", user.Id)
	return nil, nil
	// variables := map[string]interface{}{
	// 	"id": user.Id,
	// }
	// response, err := r.CassandraClient.Query(cassandra.QueryUserByIdGql, variables)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("response: %v\n", response)
	// me, err := cassandra.GetUserFromQueryUserByIdResponse(response)
	// if err != nil {
	// 	return nil, err
	// } else if me == nil {
	// 	err = errors.New("not found me")
	// 	return nil, err
	// }
	// return me, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, idOrName string) ([]*model.User, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v query idOrName: %s\n", user.Id, idOrName)
	return nil, nil
	// variables := map[string]interface{}{
	// 	"id": idOrName,
	// }
	// response, err := r.CassandraClient.Query(cassandra.QueryUserByIdGql, variables)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("query user by id response: %v\n", response)

	// mapUsers := make(map[string]*model.User)
	// queriedByIdUser, err := cassandra.GetUserFromQueryUserByIdResponse(response)
	// if err != nil {
	// 	return nil, err
	// } else if queriedByIdUser != nil {
	// 	mapUsers[queriedByIdUser.ID] = queriedByIdUser
	// }

	// variables = map[string]interface{}{
	// 	"name": idOrName,
	// }
	// for {
	// 	response, err = r.CassandraClient.Query(cassandra.QueryUserByNameGql, variables)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	fmt.Printf("query user by name response: %v\n", response)

	// 	queriedUsers, pageState, getErr := cassandra.GetUserFromQueryUserByNameResponse(response)
	// 	if getErr != nil {
	// 		err = getErr
	// 		return nil, err
	// 	}
	// 	for id, user := range queriedUsers {
	// 		mapUsers[id] = user
	// 	}
	// 	if pageState == "" {
	// 		break
	// 	} else {
	// 		variables["pageState"] = pageState
	// 	}
	// }

	// users := maps.Values(mapUsers)
	// return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
