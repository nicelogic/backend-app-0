package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"message/graph/generated"
	"message/graph/model"
)

// CreateChat is the resolver for the createChat field.
func (r *mutationResolver) CreateChat(ctx context.Context, memberIds []string, name string) (*model.Chat, error) {
	panic(fmt.Errorf("not implemented: CreateChat - createChat"))
}

// DeleteChat is the resolver for the deleteChat field.
func (r *mutationResolver) DeleteChat(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteChat - deleteChat"))
}

// GetChats is the resolver for the getChats field.
func (r *queryResolver) GetChats(ctx context.Context) ([]*model.Chat, error) {
	panic(fmt.Errorf("not implemented: GetChats - getChats"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

