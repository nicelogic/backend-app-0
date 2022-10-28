package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/graph/generated"
	"contacts/graph/model"
	"context"
	"fmt"
)

// ApplyAddContacts is the resolver for the applyAddContacts field.
func (r *mutationResolver) ApplyAddContacts(ctx context.Context, input model.ApplyAddContactsInput) (string, error) {
	panic(fmt.Errorf("not implemented: ApplyAddContacts - applyAddContacts"))
}

// ReplyAddContacts is the resolver for the replyAddContacts field.
func (r *mutationResolver) ReplyAddContacts(ctx context.Context, input model.ReplyAddContactsInput) (string, error) {
	panic(fmt.Errorf("not implemented: ReplyAddContacts - replyAddContacts"))
}

// RemoveContacts is the resolver for the removeContacts field.
func (r *mutationResolver) RemoveContacts(ctx context.Context, contactsID string) (string, error) {
	panic(fmt.Errorf("not implemented: RemoveContacts - removeContacts"))
}

// AddContactsApply is the resolver for the addContactsApply field.
func (r *queryResolver) AddContactsApply(ctx context.Context) ([]*model.AddContactsApply, error) {
	panic(fmt.Errorf("not implemented: AddContactsApply - addContactsApply"))
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, first int, after string) (*model.ContactsConnection, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
