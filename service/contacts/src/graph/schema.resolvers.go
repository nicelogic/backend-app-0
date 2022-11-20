package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/graph/generated"
	"contacts/graph/model"
	"contacts/sql"
	"context"
	"fmt"

	"github.com/nicelogic/auth"
)

// RemoveContacts is the resolver for the removeContacts field.
func (r *mutationResolver) RemoveContacts(ctx context.Context, contactsID string) (bool, error) {
	panic(fmt.Errorf("not implemented: RemoveContacts - removeContacts"))
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, first int, after string) (*model.ContactsConnection, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %#v query contacts\n", user)
	return nil, err
}

// AddedMe is the resolver for the addedMe field.
func (r *queryResolver) AddedMe(ctx context.Context, userID string) (bool, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return false, err
	}
	fmt.Printf("user: %#v query user: %s did added me\n", user, userID)
	result, err := r.CrdbClient.Query(ctx, sql.QueryUserAddedMe, userID, user.Id)
	if err != nil {
		return false, err
	}
	isUserAddedMe := len(result) != 0
	return isUserAddedMe, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

