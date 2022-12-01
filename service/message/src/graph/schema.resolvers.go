package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"message/graph/generated"
	"message/graph/model"
	"message/sql"

	pgx "github.com/jackc/pgx/v4"
	"github.com/nicelogic/authutil"
)

// CreateChat is the resolver for the createChat field.
func (r *mutationResolver) CreateChat(ctx context.Context, memberIds []string, name *string) (*model.Chat, error) {
	user, err := authutil.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("user: %#v crate chat, name(%v), members(%v)\n", user, name, memberIds)

	if len(memberIds) == 2 {
		sameMembersChatRow := r.CrdbClient.Pool.QueryRow(ctx, sql.QuerySameMembersP2pChatWhetherExist, memberIds, len(memberIds))
		var id string
		var memberIds []string
		var name, last_message *string
		err = sameMembersChatRow.Scan(&id, &memberIds, &name, &last_message)
		if err == pgx.ErrNoRows {
			log.Printf("no same members chat, can create chat\n")
		} else if err != nil {
			log.Printf("query row err: %v\n", err)
			return nil, err
		}
		log.Printf("has same members p2p chat, just return this chat")
		members := []*model.User{}
		for _, memberId := range memberIds {
			members = append(members, &model.User{ID: memberId})
		}
		return &model.Chat{
			ID:          id,
			Members:     members,
			Name:        name,
			LastMessage: nil,
		}, nil
	}

	return nil, err
}

// DeleteChat is the resolver for the deleteChat field.
func (r *mutationResolver) DeleteChat(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteChat - deleteChat"))
}

// GetChats is the resolver for the getChats field.
func (r *queryResolver) GetChats(ctx context.Context) ([]*model.Chat, error) {
	panic(fmt.Errorf("not implemented: GetChats - getChats"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
