package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"user/graph/generated"
	"user/graph/model"

	"github.com/nicelogic/auth"
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


func (r *mutationResolver) UpdateUser(ctx context.Context, changes map[string]interface{}) (user *model.User, err error) {
	requestUser, err := auth.GetUser(ctx)
	if err != nil{
		return
	}
	user = &model.User{}
	user.ID = requestUser.Id
	user.Name = "test"
	user.Signature = "well"


	return 
}

func (r *queryResolver) Me(ctx context.Context) (user *model.User, err error) {
	return 
}

func (r *queryResolver) User(ctx context.Context, idOrName string) (user *model.User, err error) {
	return 
}
