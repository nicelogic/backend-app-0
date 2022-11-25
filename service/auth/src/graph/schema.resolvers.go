package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"auth/graph/generated"
	"auth/graph/model"
	"context"
	"fmt"
)

// SignUpByUserName is the resolver for the signUpByUserName field.
func (r *mutationResolver) SignUpByUserName(ctx context.Context, userName string, pwd string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: SignUpByUserName - signUpByUserName"))
}

// SignInByUserName is the resolver for the signInByUserName field.
func (r *queryResolver) SignInByUserName(ctx context.Context, userName string, pwd string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: SignInByUserName - signInByUserName"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
