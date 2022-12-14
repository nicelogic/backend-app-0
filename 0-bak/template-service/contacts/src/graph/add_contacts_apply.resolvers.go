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
func (r *mutationResolver) ApplyAddContacts(ctx context.Context, input model.ApplyAddContactsInput) (bool, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return false, err
	}
	fmt.Printf("user: %#v apply add contacts: %#v\n", user, input)

	variables := map[string]interface{}{
		"user_id": input.ContactsID,
		"contacts_id": user.Id,
	}
	response, err := r.CassandraClient.Query(cassandra.User_contacts_record, variables)
	if err != nil {
		return false, err
	}
	fmt.Printf("user contacts record response: %v\n", response)
	_, jsonValue, err := r.CassandraClient.QueryResponse(response)
	if err != nil {
		return false, err
	}
	if string(jsonValue) != "[]" {
		fmt.Printf("user: %v want add contacts: %v, but contacts already added u, needn't apply\n", user.Id, input.ContactsID)
		err = errors.New(ContactsAlreadyAddedU)
		return false, err
	}

	variables = map[string]interface{}{
		"user_id":                user.Id,
		"contacts_id":            input.ContactsID,
		"remark_name":            input.RemarkName,
		"message":                input.Message,
		"update_time":            time.Now().Format(time.RFC3339),
		"add_contacts_apply_ttl": 604800,
	}
	response, err = r.CassandraClient.Mutation(cassandra.UpdateAddContactsApplyGql, variables)
	if err != nil {
		return false, err
	}
	fmt.Println(response)

	mutations := []string{
		"updatecontacts_by_remark_name",
		"updatecontacts",
		"updateadd_contacts_apply",
	}
	_, err = r.CassandraClient.BatchMutationResponse(response, mutations)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ReplyAddContacts is the resolver for the replyAddContacts field.
func (r *mutationResolver) ReplyAddContacts(ctx context.Context, input model.ReplyAddContactsInput) (bool, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return false, err
	}
	fmt.Printf("user: %#v reply add contacts apply\n", user)

	return true, err
}

// AddContactsApply is the resolver for the addContactsApply field.
func (r *queryResolver) AddContactsApply(ctx context.Context, first *int, after *string) (*model.AddContactsApplyConnection, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %#v query add contacts apply\n", user)

	variables := map[string]interface{}{
		"user_id": user.Id,
		"first":   first,
		"after":   after,
	}
	response, err := r.CassandraClient.Query(cassandra.AddContactsApplyGql, variables)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)

	pageState, jsonValue, err := r.CassandraClient.QueryResponse(response)
	if err != nil {
		return nil, err
	}
	addContactsApplys := make([]model.AddContactsApply, 0)
	err = json.Unmarshal(jsonValue, &addContactsApplys)
	if err != nil {
		return nil, err
	}
	addContactsApplyConnection := &model.AddContactsApplyConnection{}
	addContactsApplyConnection.TotalCount = len(addContactsApplys)
	for _, apply := range addContactsApplys {
		apply := apply
		edge := &model.AddContactsApplyEdge{}
		edge.Node = &apply
		addContactsApplyConnection.Edges = append(addContactsApplyConnection.Edges, edge)
	}
	addContactsApplyConnection.PageInfo = &model.AddContactsApplyEdgePageInfo{}
	addContactsApplyConnection.PageInfo.EndCursor = pageState
	addContactsApplyConnection.PageInfo.HasNextPage = pageState != nil

	return addContactsApplyConnection, err
}
