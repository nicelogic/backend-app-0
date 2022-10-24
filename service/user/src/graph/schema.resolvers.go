package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
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

func (r *mutationResolver) UpdateUser2(ctx context.Context, name *string, signature *string) (responseUser *model.User, err error) {
	return
}

func (r *mutationResolver) UpdateUser(ctx context.Context, changes map[string]interface{}) (responseUser *model.User, err error) {
	user, err := auth.GetUser(ctx)
	if err != nil{
		return
	}
	fmt.Printf("user: %v update\n", user.Id)

	updatedUser := &model.User{}
	var metadata mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &metadata,
		Result:   &updatedUser,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return
	}
	err = decoder.Decode(changes)
	if err != nil{
		return
	}
	fmt.Printf("changes: %v\n", changes)
	fmt.Printf("success decoded: %v\n", metadata.Keys)
	for _, value := range metadata.Keys{
		fmt.Printf("%s change to: %v\n", value, changes["Signature"])
	}

	const gql = `mutation updateUser($user_id: String!, $update_time: Timestamp!, $name: String, $signature: String) {
		updateUserName: updateuser(value: {
										  user_id: $user_id
						  name:$name , 
										  signature: $signature
						  update_time: $update_time
						},
									  ifExists: false,
						
						)
		{
				applied,
				accepted,
				value {
				  user_id,
				  name,
				  signature,
				  update_time
	  
				}
			  }
	  }`
	variables := map[string]interface{}{
		"user_id": user.Id,
  		"update_time": "2022-10-22T04:03:18.879Z",
		"name": "hi",
		"signature": "hi",
	}
	response, err := r.CassandraClient.Mutation(gql, variables)
	if err != nil{
		return
	}
	fmt.Printf("%v", response)

	responseUser = &model.User{}
	err = mapstructure.Decode(response, &responseUser)
	if err != nil{
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
