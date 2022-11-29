package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"message/graph/generated"
	"message/graph/model"
)

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, chatID string, message string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: CreateMessage - createMessage"))
}

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context, chatID string) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: GetMessages - getMessages"))
}

// NewMessageReceived is the resolver for the newMessageReceived field.
func (r *subscriptionResolver) NewMessageReceived(ctx context.Context, userID string) (<-chan *model.NewChatMessage, error) {
	panic(fmt.Errorf("not implemented: NewMessageReceived - newMessageReceived"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
