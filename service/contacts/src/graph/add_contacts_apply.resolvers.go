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

	"github.com/jackc/pgx/v4"
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
	err = r.CrdbClient.Pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql.UpsertAddContactsApply, user.Id, input.ContactsID, input.Message, updateTime)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, sql.UpsertContacts, user.Id, input.ContactsID, input.RemarkName, updateTime)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Printf("transcation func err: %v\n", err)
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

	return nil, err
}
