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
	fmt.Println(response)

	addContactsApply := &model.AddContactsApply{}
	applied, jsonValue, err := r.CassandraClient.MutationResponse(response)
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
	if err != nil{
		return nil, err
	}
	addContactsApplys := make([]model.AddContactsApply, 0)
	err = json.Unmarshal(jsonValue, &addContactsApplys)
	if err != nil{
		return nil, err
	}
	addContactsApplyConnection := &model.AddContactsApplyConnection{}
	addContactsApplyConnection.TotalCount = len(addContactsApplys)
	for _, apply := range addContactsApplys {
		edge := &model.AddContactsApplyEdge{}
		edge.Node = &apply
		addContactsApplyConnection.Edges = append(addContactsApplyConnection.Edges, edge)
	}
	addContactsApplyConnection.PageInfo = &model.AddContactsApplyEdgePageInfo{}
	addContactsApplyConnection.PageInfo.EndCursor = pageState
	addContactsApplyConnection.PageInfo.HasNextPage = pageState != nil

	return addContactsApplyConnection, err
}
