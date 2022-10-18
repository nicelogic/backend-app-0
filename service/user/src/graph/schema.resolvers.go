package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"user/graph/generated"
	"user/graph/model"
)

type contextKey struct {
	name string
}
var userCtxKey = &contextKey{"user"}

type User struct {
	Id string
}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			user := &User{Id: "wzh"}
			fmt.Println("has user")
			ctx := context.WithValue(request.Context(), userCtxKey, user)
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	user := ForContext(ctx)
	fmt.Printf("userId: %v", user.Id)

	return []*model.Todo{}, nil
	//panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
