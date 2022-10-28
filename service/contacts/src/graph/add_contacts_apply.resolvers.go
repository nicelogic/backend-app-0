package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/cassandra"
	"contacts/graph/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/nicelogic/auth"
)

// ApplyAddContacts is the resolver for the applyAddContacts field.
func (r *mutationResolver) ApplyAddContacts(ctx context.Context, input model.ApplyAddContactsInput) (*model.AddContactsApply, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	addContactsApply := &model.AddContactsApply{}
	applied, jsonValue, err := r.CassandraClient.GetValue(response)
	if err != nil {
		return nil, err
	}
	if !applied {
		err = errors.New("cassandra not applied")
		return nil, err
	}
	err = json.Unmarshal(jsonValue, addContactsApply)
	if err != nil {
		return nil, err
	}
	return addContactsApply, nil
}

// ReplyAddContacts is the resolver for the replyAddContacts field.
func (r *mutationResolver) ReplyAddContacts(ctx context.Context, input model.ReplyAddContactsInput) (bool, error) {
	panic(fmt.Errorf("not implemented: ReplyAddContacts - replyAddContacts"))
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
