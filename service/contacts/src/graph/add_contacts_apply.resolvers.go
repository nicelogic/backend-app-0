package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"contacts/graph/model"
	"contacts/sql"
	"context"
	"fmt"
	"log"
	"time"

	pgx "github.com/jackc/pgx/v4"
	"github.com/nicelogic/auth"
)

// ApplyAddContacts is the resolver for the applyAddContacts field.
func (r *mutationResolver) ApplyAddContacts(ctx context.Context, input model.ApplyAddContactsInput) (bool, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return false, err
	}
	log.Printf("user: %#v apply add contacts: %#v\n", user, input)

	updateTime := time.Now().Format(time.RFC3339)
	_, err = r.CrdbClient.Pool.Exec(ctx, sql.UpsertContacts, user.Id, input.ContactsID, input.RemarkName, updateTime)
	if err != nil {
		log.Printf("sql upsert contacts err: %v\n", err)
		return false, err
	}
	err = r.CrdbClient.Pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql.QueryUserAddedMe, input.ContactsID, user.Id)
		if err != nil && rows.Err() != nil {
			return rows.Err()
		}
		defer rows.Close()
		if rows.Next() {
			return fmt.Errorf(ContactsAddedMe)
		}
		_, err = tx.Exec(ctx, sql.UpsertAddContactsApply, user.Id, input.ContactsID, input.Message, updateTime)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Printf("transcation func err: %v\n", err)
		return false, err
	}
	log.Printf("apply success")
	return true, nil
}

// ReplyAddContacts is the resolver for the replyAddContacts field.
func (r *mutationResolver) ReplyAddContacts(ctx context.Context, input model.ReplyAddContactsInput) (bool, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return false, err
	}
	log.Printf("user: %#v reply add contacts apply, isAgree: %v\n", user, input.IsAgree)
	if !input.IsAgree {
		_, err = r.CrdbClient.Pool.Exec(ctx, sql.DeleteAddContactsApply, user.Id, input.ContactsID)
		if err != nil {
			log.Printf("delete addContactsApply fail, err: %v\n", err)
			return false, err
		}
	} else {
		updateTime := time.Now().Format(time.RFC3339)
		err = r.CrdbClient.Pool.BeginFunc(ctx, func(tx pgx.Tx) error {
			_, err = r.CrdbClient.Pool.Exec(ctx, sql.UpsertContacts, user.Id, input.ContactsID, input.RemarkName, updateTime)
			if err != nil {
				log.Printf("sql upsert contacts err: %v\n", err)
				return err
			}
			_, err = r.CrdbClient.Pool.Exec(ctx, sql.DeleteAddContactsApply, user.Id, input.ContactsID)
			if err != nil {
				log.Printf("delete addContactsApply err: %v\n", err)
				return err
			}
			return err
		})
		if err != nil {
			log.Printf("transcation func err: %v\n", err)
			return false, err
		}
	}
	log.Printf("reply success")
	return true, err
}

// AddContactsApply is the resolver for the addContactsApply field.
func (r *queryResolver) AddContactsApply(ctx context.Context, first *int, after *string) (*model.AddContactsApplyConnection, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %#v query add contacts apply\n", user)

	cursor := time.Now().Format(time.RFC3339)
	if after != nil {
		cursor = *after
	}
	addContactsApplys, err := r.CrdbClient.Query(ctx, sql.QueryAddContactsApply, user.Id, cursor, *first)
	if err != nil {
		fmt.Printf("query err: %v\n", err)
		return nil, err
	}
	addContactsApplyConnection := &model.AddContactsApplyConnection{}
	addContactsApplyConnection.TotalCount = len(addContactsApplys)
	endCursor := ""
	for _, addContactsApply := range addContactsApplys {
		fmt.Printf("apply: %#v\n", addContactsApply)
		addContactsApply := addContactsApply.([]any)
		edge := &model.AddContactsApplyEdge{}
		apply := model.AddContactsApply{}
		apply.ContactsID = user.Id
		apply.UserID = addContactsApply[0].(string)
		apply.Message = addContactsApply[1].(string)
		apply.UpdateTime = addContactsApply[2].(time.Time).Format(time.RFC3339)
		edge.Node = &apply
		addContactsApplyConnection.Edges = append(addContactsApplyConnection.Edges, edge)
		endCursor = apply.UpdateTime
	}
	addContactsApplyConnection.PageInfo = &model.AddContactsApplyEdgePageInfo{}
	addContactsApplyConnection.PageInfo.EndCursor = &endCursor
	addContactsApplyConnection.PageInfo.HasNextPage = addContactsApplyConnection.TotalCount == *first
	return addContactsApplyConnection, err
}
