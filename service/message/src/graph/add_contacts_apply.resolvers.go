package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	contactserror "contacts/graph/error"
	"contacts/graph/generated"
	"contacts/graph/model"
	"contacts/sql"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
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
			return fmt.Errorf(contactserror.ContactsAddedMe)
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

	ntf := &model.AddContactsApplyNtf{
		UserID:   user.Id,
		UserName: input.UserName,
	}
	msgId, err := r.PulsarClient.Send(ctx, ntf)
	if err != nil{
		log.Printf("pulsar send ntf: %v, err: %v\n", ntf, err)
	}
	log.Printf("pulsar send ntf: %v, success, msgId: %v\n", ntf, msgId)
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

	updateTime := time.Now().Format(time.RFC3339)
	applyAddMeUserId := ""
	if after != nil {
		decodeAfter, _ := base64.StdEncoding.DecodeString(*after)
		args := strings.Split(string(decodeAfter), "|")
		updateTime = args[0]
		applyAddMeUserId = args[1]
	}
	fmt.Printf("after: updateTime: %s, contactsId: %s\n", updateTime, applyAddMeUserId)
	addContactsApplys, err := r.CrdbClient.Query(ctx, sql.QueryAddContactsApply, user.Id, updateTime, applyAddMeUserId, *first)
	if err != nil {
		fmt.Printf("query err: %v\n", err)
		return nil, err
	}
	addContactsApplyConnection := &model.AddContactsApplyConnection{}
	addContactsApplyConnection.TotalCount = len(addContactsApplys)
	for _, addContactsApply := range addContactsApplys {
		fmt.Printf("apply: %#v\n", addContactsApply)
		addContactsApply := addContactsApply.([]any)
		apply := model.AddContactsApply{}
		apply.ContactsID = user.Id
		apply.UserID = addContactsApply[0].(string)
		apply.Message = addContactsApply[1].(string)
		apply.UpdateTime = addContactsApply[2].(time.Time).Format(time.RFC3339)
		edge := &model.AddContactsApplyEdge{}
		edge.Node = &apply
		addContactsApplyConnection.Edges = append(addContactsApplyConnection.Edges, edge)
	}
	addContactsApplyConnection.PageInfo = &model.AddContactsApplyEdgePageInfo{}
	if addContactsApplyConnection.TotalCount != 0 {
		lastNode := addContactsApplyConnection.Edges[len(addContactsApplyConnection.Edges)-1].Node
		lastUpdateTime := lastNode.UpdateTime
		lastContactsId := lastNode.UserID
		endCursor := lastUpdateTime + "|" + lastContactsId
		base64EndCursor := base64.StdEncoding.EncodeToString([]byte(endCursor))
		addContactsApplyConnection.PageInfo.EndCursor = &base64EndCursor
	}
	addContactsApplyConnection.PageInfo.HasNextPage = addContactsApplyConnection.TotalCount == *first
	return addContactsApplyConnection, err
}

// AddContactsApplyReceived is the resolver for the addContactsApplyReceived field.
func (r *subscriptionResolver) AddContactsApplyReceived(ctx context.Context, token string) (<-chan *model.AddContactsApplyNtf, error) {
	ch := make(chan *model.AddContactsApplyNtf)
	go func(token string) {
		//subscription check user by payload, because playground not work
		//check flutter client test whether ok, then do optimize
		user, err := auth.UserFromJwt(token)
		if err != nil {
			log.Printf("subscription token err: %v\n", err)
			return
		}
		log.Printf("user: %v subscribe addContactsApply ntf\n", user)

		consumer, err := r.PulsarClient.Client.Subscribe(pulsar.ConsumerOptions{
			Topic:            r.Config.Pulsar_topic,
			SubscriptionName: user.Id,
			Type: pulsar.Failover,
		})
		if err != nil {
			log.Printf("pulsar subscribe err: %v\n", err)
			return
		}
		defer consumer.Close()
		for {
			ntf := &model.AddContactsApplyNtf{}
			err = r.PulsarClient.Receive(ctx, consumer, ntf)
			if err != nil{
				log.Printf("pulsar receive err: %v\n", err)
				continue
			}
			select {
			case ch <- ntf:
				log.Printf("send addContactsApplyNtf: %v\n", ntf)
			case <-ctx.Done():
				log.Println("ctx done")
				return
			}

		}
	}(token)

	return ch, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
