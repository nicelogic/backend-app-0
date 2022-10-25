package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	userConfig "user/config"
	"user/graph"
	"user/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/nicelogic/auth"
	"github.com/nicelogic/cassandra"
	"github.com/nicelogic/common_error"
	"github.com/nicelogic/config"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {

	client := cassandra.Client{}
	client.Init("app_0_user")
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{ CassandraClient: &client}}))
	server.SetRecoverFunc(func(ctx context.Context, panicErr interface{}) error {
		log.Print(panicErr)
		err := &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    common_error.ServerInternalError,
			Extensions: map[string]interface{}{},
		}
		return err
	})
	server.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		fmt.Printf("error: %v\n", e)
		err := graphql.DefaultErrorPresenter(ctx, e)
		var jwtError *jwt.ValidationError
		hasJwtError := errors.As(e, &jwtError)
		switch{
		case hasJwtError && jwtError.Errors == jwt.ValidationErrorExpired:
			err.Message = common_error.TokenExpired
		case hasJwtError:
			err.Message = common_error.TokenInvalid
		case err.Message == UserNotExist:
			break
		default:
			err.Message = common_error.ServerInternalError
		}
		return err
	})

	userConfig := userConfig.Config{Path: "/", Listen_address: ":80"}
	config.Init("/etc/app-0/config-user/config-user.yml", &userConfig)
	path := userConfig.Path
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle(path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://%s%s for GraphQL playground", userConfig.Listen_address, userConfig.Path)
	log.Fatal(http.ListenAndServe(userConfig.Listen_address, router))
}
