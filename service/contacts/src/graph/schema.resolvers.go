package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/cassandra"
	"contacts/graph/generated"
	"contacts/graph/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nicelogic/auth"
)

// ApplyAddContacts is the resolver for the applyAddContacts field.
func (r *mutationResolver) ApplyAddContacts(ctx context.Context, input model.ApplyAddContactsInput) (string, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return "", err
	}
	fmt.Printf("user: %#v apply add contacts: %#v\n", user, input)

	appyId := user.Id + ">" + input.ContactsID
	variables := map[string]interface{}{
		"id":          appyId,
		"user_id":     user.Id,
		"contacts_id": input.ContactsID,
		"remark_name": input.RemarkName,
		"update_time": time.Now().Format(time.RFC3339),
	}
	response, err := r.CassandraClient.Mutation(cassandra.UpdateAddContactsApplyGql, variables)
	if err != nil {
		return "", err
	}
	response = response["response"].(map[string]interface{})
	applied := response["applied"].(bool)
	if !applied {
		err = errors.New("cassandra not applied")
		return "", err
	}
	return appyId, nil
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
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %#v query add contacts apply\n", user)
	return nil, err
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
