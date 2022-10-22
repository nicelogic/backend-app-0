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


func (r *mutationResolver) UpdateUser(ctx context.Context, changes map[string]interface{}) (updatedUser *model.User, err error) {
	user, err := auth.GetUser(ctx)
	if err != nil{
		return
	}
	fmt.Printf("user: %v update\n", user.Id)

	updatedUser = &model.User{}
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
	fmt.Printf("success decoded: %v\n", metadata.Keys)


	return 
}

func (r *queryResolver) Me(ctx context.Context) (user *model.User, err error) {
	return 
}

func (r *queryResolver) User(ctx context.Context, idOrName string) (user *model.User, err error) {
	return 
}
