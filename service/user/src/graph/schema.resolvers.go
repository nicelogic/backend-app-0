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

	"github.com/mitchellh/mapstructure"
	"github.com/nicelogic/auth"
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
	fmt.Printf("%v\n", response)

	response, ok := response["updateUser"].(map[string]interface{})
	if !ok {
		err = errors.New("response[updateUser]'s type is not map[string]interface{}")
		return
	}
	applied, ok := response["applied"].(bool)
	if !ok {
		err = errors.New("response[updateUser][applied]'s type is not bool")
		return
	}
	if !applied {
		err = errors.New("cassandra not applied")
		return
	}
	value, ok := response["value"].(map[string]interface{})
	if !ok {
		err = errors.New("response[updateUser][value]'s type is not map[string]interface{}")
		return
	}
	responseUser = &model.User{}
	err = mapstructure.Decode(value, &responseUser)
	if err != nil {
		return
	}

	return
}

func (r *queryResolver) Me(ctx context.Context) (user *model.User, err error) {
	return
}

func (r *queryResolver) User(ctx context.Context, idOrName string) (user *model.User, err error) {
	return
}
