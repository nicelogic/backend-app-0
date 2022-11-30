package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/graph/generated"
	"contacts/graph/model"
	"contacts/sql"
	"context"
	"encoding/base64"
	"log"
	"strings"

	"github.com/nicelogic/authutil"
)

// RemoveContacts is the resolver for the removeContacts field.
func (r *mutationResolver) RemoveContacts(ctx context.Context, contactsID string) (bool, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return false, err
	}
	log.Printf("user: %#v remove contacts: %s\n", user, contactsID)
	_, err = r.CrdbClient.Pool.Exec(ctx, sql.DeleteContacts, user.Id, contactsID)
	if err != nil {
		log.Printf("delete contacts err: %v\n", err)
		return false, err
	}
	return true, err
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, first *int, after *string) (*model.ContactsConnection, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("user: %#v query contacts\n", user)
	remarkName := ""
	contactsId := ""
	if after != nil {
		decodeAfter, _ := base64.StdEncoding.DecodeString(*after)
		args := strings.Split(string(decodeAfter), "|")
		remarkName = args[0]
		contactsId = args[1]
	}
	contactsSlice, err := r.CrdbClient.Query(ctx, sql.QueryContacts, user.Id, remarkName, contactsId, first)
	if err != nil {
		log.Printf("query contacts err: %v\n", err)
		return nil, err
	}
	contactsConnection := &model.ContactsConnection{}
	contactsConnection.TotalCount = len(contactsSlice)
	for _, plainContacts := range contactsSlice {
		plainContacts := plainContacts.([]any)
		contacts := model.Contacts{}
		contacts.ID = plainContacts[0].(string)
		contacts.RemarkName = plainContacts[1].(string)
		edge := &model.Edge{}
		edge.Node = &contacts
		contactsConnection.Edges = append(contactsConnection.Edges, edge)
	}
	contactsConnection.PageInfo = &model.PageInfo{}
	if contactsConnection.TotalCount != 0 {
		lastNode := contactsConnection.Edges[len(contactsConnection.Edges)-1].Node
		lastRemarkName := lastNode.RemarkName
		lastContactsId := lastNode.ID
		endCursor := lastRemarkName + "|" + lastContactsId
		base64EndCursor := base64.StdEncoding.EncodeToString([]byte(endCursor))
		contactsConnection.PageInfo.EndCursor = &base64EndCursor
	}
	contactsConnection.PageInfo.HasNextPage = contactsConnection.TotalCount == *first
	return contactsConnection, err
}

// AddedMe is the resolver for the addedMe field.
func (r *queryResolver) AddedMe(ctx context.Context, userID string) (bool, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return false, err
	}
	log.Printf("user: %#v query user: %s did added me\n", user, userID)
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
