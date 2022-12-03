package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"message/graph/generated"
	"message/graph/model"
	"message/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v4"
	"github.com/nicelogic/authutil"
)

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, chatID string, message string) (*model.Message, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("user(%#v) create message in chat(%s)\n", user, chatID)

	messageId := uuid.New().String()
	messageTime := time.Now().Format(time.RFC3339)
	err = r.CrdbClient.Pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := r.CrdbClient.Pool.Exec(ctx, sql.InsertMessage, messageId, chatID, message, user.Id)
		if err != nil {
			log.Printf("in transcation create message , insert message err(%v)\n", err)
			return err
		}
		_, err = r.CrdbClient.Pool.Exec(ctx, sql.UpdateChatLastMessage, message, messageTime, chatID)
		if err != nil {
			log.Printf("in transcation create message , insert message err(%v)\n", err)
			return err
		}
		return err
	})
	if err != nil {
		log.Printf("create message transcation err(%v)\n", err)
		return nil, err
	}
	createdMessage := model.Message{
		ID:      messageId,
		Content: message,
		Sender: &model.User{
			ID: user.Id,
		},
		CreateTime: messageTime,
	}
	//public new message to chat members
	return &createdMessage, err
}

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context, chatID string, first *int, after *string) (*model.MessageConnection, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("user(%#v) get messages in chat(%s)\n", user, chatID)

	messageCreateTime := time.Now().Format(time.RFC3339)
	messageId := ""
	if after != nil {
		decodeAfter, _ := base64.StdEncoding.DecodeString(*after)
		args := strings.Split(string(decodeAfter), ",")
		messageCreateTime = args[1]
		messageId = args[2]
	}
	log.Printf("after: createTime: %s, messageId: %s\n", messageCreateTime, messageId)
	messagesValues, err := r.CrdbClient.Query(ctx, sql.QueryMessages, chatID, messageCreateTime, messageId, *first)
	if err != nil {
		log.Printf("query err: %v\n", err)
		return nil, err
	}
	messageConnection := &model.MessageConnection{}
	for _, messageValues := range messagesValues {
		fmt.Printf("messageVlues: %#v\n", messageValues)
		messageValues := messageValues.([]any)
		message := model.Message{
			ID: messageValues[0].(string),
			Content: messageValues[1].(string),	
			Sender: &model.User{
				ID: messageValues[2].(string),
			},
			CreateTime: messageValues[3].(time.Time).Format(time.RFC3339),
		}
		edge := &model.MessageEdge{}
		edge.Node = &message
		messageConnection.Edges = append(messageConnection.Edges, edge)
	}
	messageConnection.TotalCount = len(messageConnection.Edges)
	messageConnection.PageInfo = &model.MessagePageInfo{}
	if messageConnection.TotalCount != 0 {
		lastNode := messageConnection.Edges[len(messageConnection.Edges)-1].Node
		lastMessageTime := lastNode.CreateTime
		lastMessageId := lastNode.ID
		endCursor := lastMessageTime + "," + lastMessageId
		base64EndCursor := base64.StdEncoding.EncodeToString([]byte(endCursor))
		messageConnection.PageInfo.EndCursor = &base64EndCursor
	}
	messageConnection.PageInfo.HasNextPage = messageConnection.TotalCount == *first
	return messageConnection, nil
}

// NewMessageReceived is the resolver for the newMessageReceived field.
func (r *subscriptionResolver) NewMessageReceived(ctx context.Context, userID string) (<-chan *model.NewChatMessage, error) {
	panic(fmt.Errorf("not implemented: NewMessageReceived - newMessageReceived"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
