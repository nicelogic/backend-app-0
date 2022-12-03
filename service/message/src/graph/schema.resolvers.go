package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"message/constant"
	"message/graph/generated"
	"message/graph/model"
	"message/sql"
	"sort"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v4"
	"github.com/nicelogic/authutil"
)

// CreateChat is the resolver for the createChat field.
/*
	create p2p chat need check whether p2p chat exist
	but needn't transaction
	because the id: "userid_low" | "userid_high" decide the p2p chat's id
	create p2p chat at the same time, created after will fail
	if group, It is to create multiple, do not judge whether there are the same members
*/
func (r *mutationResolver) CreateChat(ctx context.Context, memberIds []string, name *string) (*model.Chat, error) {
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
		sameMembersChatRow := r.CrdbClient.Pool.QueryRow(ctx, sql.QueryChat, user.Id, chatId)
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
	createdChatRow := r.CrdbClient.Pool.QueryRow(ctx, sql.QueryChat, user.Id, chatId)
	createdChat, err := sql.ChatRowToChatModel(createdChatRow)
	if err != nil {
		log.Printf("query new created chat err(%v)\n", err)
		return nil, err
	}
	log.Printf("find chat, return that chat\n")
	return createdChat, nil
}

// UpdateChatSetting is the resolver for the updateChatSetting field.
func (r *mutationResolver) UpdateChatSetting(ctx context.Context, id string, setting string) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateChatSetting - updateChatSetting"))
}

// NotShowChat is the resolver for the notShowChat field.
func (r *mutationResolver) NotShowChat(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: NotShowChat - notShowChat"))
}

// AddGroupChatMembers is the resolver for the addGroupChatMembers field.
func (r *mutationResolver) AddGroupChatMembers(ctx context.Context, id string, memberIds []string) (bool, error) {
	panic(fmt.Errorf("not implemented: AddGroupChatMembers - addGroupChatMembers"))
}

// RemoveGroupChatMembers is the resolver for the removeGroupChatMembers field.
func (r *mutationResolver) RemoveGroupChatMembers(ctx context.Context, id string, memberIds []string) (bool, error) {
	panic(fmt.Errorf("not implemented: RemoveGroupChatMembers - removeGroupChatMembers"))
}

// DeleteGroupChat is the resolver for the deleteGroupChat field.
func (r *mutationResolver) DeleteGroupChat(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteGroupChat - deleteGroupChat"))
}

// GetChats is the resolver for the getChats field.
func (r *queryResolver) GetChats(ctx context.Context, first *int, after *string) (*model.ChatConnection, error) {
	panic(fmt.Errorf("not implemented: GetChats - getChats"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
