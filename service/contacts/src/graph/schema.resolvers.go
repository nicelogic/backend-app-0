package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/graph/generated"
	"contacts/graph/model"
	"context"
	"fmt"
)

// AddContacts is the resolver for the addContacts field.
func (r *mutationResolver) AddContacts(ctx context.Context, contactsID string) (string, error) {
	panic(fmt.Errorf("not implemented: AddContacts - addContacts"))
}

// RemoveContacts is the resolver for the removeContacts field.
func (r *mutationResolver) RemoveContacts(ctx context.Context, contactsID string) (string, error) {
	panic(fmt.Errorf("not implemented: RemoveContacts - removeContacts"))
}

// PaginationContacts is the resolver for the paginationContacts field.
func (r *queryResolver) PaginationContacts(ctx context.Context, userID string, first int, after string) (*model.PaginationContacts, error) {
	panic(fmt.Errorf("not implemented: PaginationContacts - paginationContacts"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
