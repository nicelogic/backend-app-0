package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"message/graph/generated"
	"message/graph/model"
	"message/sql"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
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
	createdMessage := model.Message{
		ID:      messageId,
		Content: message,
		Sender: &model.User{
			ID: user.Id,
		},
		CreateTime: messageTime,
	}
	createdMessageJson, err := json.Marshal(createdMessage)
	if err != nil {
		log.Printf("json marshal created message err(%v)\n", err)
		return nil, err
	}
	err = r.CrdbClient.Pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := r.CrdbClient.Pool.Exec(ctx, sql.InsertMessage, messageId, chatID, message, user.Id)
		if err != nil {
			log.Printf("in transcation create message , insert message err(%v)\n", err)
			return err
		}
		_, err = r.CrdbClient.Pool.Exec(ctx, sql.UpdateChatLastMessage, createdMessageJson, messageTime, chatID)
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
	go func() {
		sendNewMessageNtfCtx := context.Background()
		chatRow := r.CrdbClient.Pool.QueryRow(sendNewMessageNtfCtx, sql.QueryUserChat, user.Id, chatID)
		chat, err := sql.ChatRowToChatModel(chatRow)
		if err != nil {
			log.Printf("query chat members err(%v)\n", err)
			return
		}
		chatMembers := chat.Members
		for _, member := range chatMembers {
			ntf := &model.NewMessage{
				Message: &createdMessage,
			}
			userNewMessageTopic := r.Config.Pulsar_topic + "-" + member.ID
			producer, err := r.PulsarClient.Client.CreateProducer(pulsar.ProducerOptions{
				Topic: userNewMessageTopic,
			})
			if err != nil {
				log.Printf("create producer err: %v\n", err)
				return
			}
			defer producer.Close()
			payload, err := json.Marshal(ntf)
			if err != nil {
				log.Printf("json marshal err: %v\n", err)
				return 
			}
			msgId, err := producer.Send(sendNewMessageNtfCtx, &pulsar.ProducerMessage{
				Payload: payload,
			})
			if err != nil {
				log.Printf("producer send err: %v\n", err)
				return 
			}
			log.Printf("pulsar send new message ntf(%v) to member(%v) success, msgId(%v)\n", ntf, member, msgId)
		}
	}()
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
		messageCreateTime = args[0]
		messageId = args[1]
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
		content := messageValues[1].(map[string]interface{})
		byteContent, err := json.Marshal(content)
		if err != nil {
			log.Printf("json marshal content err(%v)\n", err)
			return nil, err
		}
		message := model.Message{
			ID:      messageValues[0].(string),
			Content: string(byteContent),
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
func (r *subscriptionResolver) NewMessageReceived(ctx context.Context, token string) (<-chan *model.NewMessage, error) {
	ch := make(chan *model.NewMessage)
	go func(token string) {
		//subscription check user by payload, because playground not work
		//check flutter client test whether ok, then do optimize
		user, err := r.AuthUtil.UserFromJwt(token)
		if err != nil {
			log.Printf("subscription token err(%v)\n", err)
			return
		}
		log.Printf("user(%v) subscribe new message ntf\n", user)
		userNewMessageTopic := r.Config.Pulsar_topic + "-" + user.Id
		consumer, err := r.PulsarClient.Client.Subscribe(pulsar.ConsumerOptions{
			Topic:            userNewMessageTopic,
			SubscriptionName: user.Id,
			Type:             pulsar.Failover,
		})
		if err != nil {
			log.Printf("pulsar subscribe err(%v)\n", err)
			return
		}
		defer consumer.Close()
		for {
			ntf := &model.NewMessage{}
			err = r.PulsarClient.Receive(ctx, consumer, ntf)
			if err != nil {
				log.Printf("pulsar receive err(%v)\n", err)
				continue
			}
			select {
			case ch <- ntf:
				log.Printf("send new message ntf(%v)\n", ntf)
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
