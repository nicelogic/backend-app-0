package main

import (
	contactsConfig "contacts/config"
	"contacts/constant"
	"contacts/graph"
	"contacts/graph/generated"
	"context"
	"crdb"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/nicelogic/auth"
	"github.com/nicelogic/config"
	"github.com/nicelogic/errs"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	userConfig := contactsConfig.Config{
		ServiceName:            "contacts",
		CrdbPoolConnectionsNum: 4,
		CrdbConfigFilePath:     "/Users/bryan.wu/code/secret/config-crdb.yml",
		Path:                   "/",
		Listen_address:         ":80"}
	config.Init(constant.ConfigPath, &userConfig)
	crdbClient := crdb.Client{}
	err := crdbClient.Init(context.Background(),
		userConfig.CrdbConfigFilePath,
		userConfig.ServiceName,
		userConfig.CrdbPoolConnectionsNum)
	if err != nil {
		log.Fatalf("crdb init err: %v\n", err)
	}
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{CrdbClient: &crdbClient}}))
	server.SetRecoverFunc(func(ctx context.Context, panicErr interface{}) error {
		fmt.Printf("panic: %v\n", panicErr)
		err := &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    errs.ServerInternalError,
			Extensions: map[string]interface{}{},
		}
		return err
	})
	server.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		fmt.Printf("error: %v\n", e)
		err := graphql.DefaultErrorPresenter(ctx, e)
		var jwtError *jwt.ValidationError
		hasJwtError := errors.As(e, &jwtError)
		switch {
		case hasJwtError && jwtError.Errors == jwt.ValidationErrorExpired:
			err.Message = errs.TokenExpired
		case hasJwtError:
			err.Message = errs.TokenInvalid
		case graph.ContactsAlreadyAddedU == e.Error():
			fmt.Printf("ContactsAlreadyAddedU")
		default:
			err.Message = errs.ServerInternalError
		}
		return err
	})

	path := userConfig.Path
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle(path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://%s%s for GraphQL playground", userConfig.Listen_address, userConfig.Path)
	log.Fatal(http.ListenAndServe(userConfig.Listen_address, router))
}
