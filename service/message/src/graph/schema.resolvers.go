package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"message/constant"
	"message/graph/generated"
	"message/graph/model"
	"message/sql"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v4"
	"github.com/nicelogic/authutil"
)

// CreateChat is the resolver for the createChat field.
func (r *mutationResolver) CreateChat(ctx context.Context, memberIds []string, name *string) (*model.Chat, error) {
	/*
		create p2p chat need check whether p2p chat exist
		but needn't transaction
		because the id: "userid_low" | "userid_high" decide the p2p chat's id
		create p2p chat at the same time, created after will fail
		if group, It is to create multiple, do not judge whether there are the same members
	*/
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("user: %#v create chat, name(%v), members(%v)\n", user, name, memberIds)
	memberIds = append(memberIds, user.Id)
	chatId := uuid.New().String()
	chatType := constant.ChatTypeGroup
	if len(memberIds) == 2 {
		sort.Strings(memberIds)
		chatId = memberIds[0] + "|" + memberIds[1]
		chatType = constant.ChatTypeP2p
		sameMembersChatRow := r.CrdbClient.Pool.QueryRow(ctx, sql.QueryUserChat, user.Id, chatId)
		sameMembersChat, err := sql.ChatRowToChatModel(sameMembersChatRow)
		if err == pgx.ErrNoRows {
			log.Printf("no same memerbs p2p chat, can create new p2p chat\n")
		} else if err != nil {
			return nil, err
		} else {
			log.Printf("find same members p2p chat,just return that chat\n")
			return sameMembersChat, nil
		}
	}
	if name == nil {
		emptyString := ""
		name = &emptyString
	}
	log.Printf("begin create new chat\n")
	err = r.CrdbClient.Pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := r.CrdbClient.Pool.Exec(ctx, sql.InsertChat, chatId, chatType, memberIds, *name)
		if err != nil {
			log.Printf("in transcation create chat , insert chat err(%v)\n", err)
			return err
		}
		for _, memberId := range memberIds {
			_, err = r.CrdbClient.Pool.Exec(ctx, sql.InsertUserChat, memberId, chatId)
			if err != nil {
				log.Printf("in transaction create user chat, insert user chat err(%v)\n", err)
				return err
			}
		}
		return err
	})
	if err != nil {
		log.Printf("create chat transcation err(%v)\n", err)
		return nil, err
	}
	log.Printf("apply success")
	log.Printf("begin query new created chat\n")
	createdChatRow := r.CrdbClient.Pool.QueryRow(ctx, sql.QueryUserChat, user.Id, chatId)
	createdChat, err := sql.ChatRowToChatModel(createdChatRow)
	if err != nil {
		log.Printf("query new created chat err(%v)\n", err)
		return nil, err
	}
	log.Printf("find chat, return that chat\n")
	return createdChat, nil
}

// AddGroupChatMembers is the resolver for the addGroupChatMembers field.
func (r *mutationResolver) AddGroupChatMembers(ctx context.Context, id string, memberIds []string) (bool, error) {
	panic(fmt.Errorf("not implemented: AddGroupChatMembers - addGroupChatMembers"))
}

// RemoveGroupChatMembers is the resolver for the removeGroupChatMembers field.
func (r *mutationResolver) RemoveGroupChatMembers(ctx context.Context, id string, memberIds []string) (bool, error) {
	panic(fmt.Errorf("not implemented: RemoveGroupChatMembers - removeGroupChatMembers"))
}

// DeleteChat is the resolver for the deleteChat field.
func (r *mutationResolver) DeleteChat(ctx context.Context, id string) (bool, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return false, err
	}
	log.Printf("user(%#v) delete chat(%s)\n", user, id)
	//TODO
	//if group chat, need group chat's admin role user
	//Only when one of the parties deletes the contact, delete the chat at the same time
	_, err = r.CrdbClient.Pool.Exec(ctx, sql.DeleteChat, id)
	if err != nil {
		log.Printf("delete chat err(%v)\n", err)
		return false, err
	}
	return true, nil
}

// GetChats is the resolver for the getChats field.
func (r *queryResolver) GetChats(ctx context.Context, first *int, after *string) (*model.ChatConnection, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("user(%#v) query chats\n", user)
	priority := 100
	lastMessagTime := time.Now().Format(time.RFC3339)
	chatId := ""
	if after != nil {
		decodeAfter, _ := base64.StdEncoding.DecodeString(*after)
		args := strings.Split(string(decodeAfter), ",")
		priority, err = strconv.Atoi(args[0])
		if err != nil {
			log.Printf("cursor decode and atoi(args[0]) for priority err(%v)\n", err)
			return nil, err
		}
		lastMessagTime = args[1]
		chatId = args[2]
	}
	log.Printf("after: priority: %v, updateTime: %s, chatId: %s\n", priority, lastMessagTime, chatId)
	chatRows, err := r.CrdbClient.Pool.Query(ctx, sql.QueryChats, user.Id, priority, lastMessagTime, chatId, *first)
	if err != nil {
		log.Printf("query err: %v\n", err)
		return nil, err
	}
	defer chatRows.Close()
	chatApplyConnection := &model.ChatConnection{}
	for chatRows.Next() {
		chat, err := sql.ChatRowToChatModel(chatRows)	
		if err != nil{
			log.Printf("chat row scan err: %v\n", err)
			return nil, err
		}
		edge := &model.Edge{}
		edge.Node = chat
		chatApplyConnection.Edges = append(chatApplyConnection.Edges, edge)
	}
	if err = chatRows.Err(); err != nil {
		log.Printf("rows error(%v)\n", err)
		return nil, err
	}
	chatApplyConnection.TotalCount = len(chatApplyConnection.Edges)
	chatApplyConnection.PageInfo = &model.PageInfo{}
	if chatApplyConnection.TotalCount != 0 {
		lastNode := chatApplyConnection.Edges[len(chatApplyConnection.Edges)-1].Node
		lastNodePriority := "0"
		if lastNode.Pinned {
			lastNodePriority = "1"
		}
		lastMessageTime := time.Now().Format(time.RFC3339)
		if lastNode.LastMessageTime != nil{
			lastMessageTime = *lastNode.LastMessageTime
		}
		lastChatId := lastNode.ID
		endCursor := lastNodePriority + "," + lastMessageTime + "," + lastChatId
		base64EndCursor := base64.StdEncoding.EncodeToString([]byte(endCursor))
		chatApplyConnection.PageInfo.EndCursor = &base64EndCursor
	}
	chatApplyConnection.PageInfo.HasNextPage = chatApplyConnection.TotalCount == *first
	return chatApplyConnection, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
