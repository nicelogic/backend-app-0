package main

import (
	contactsConfig "contacts/config"
	"contacts/constant"
	"contacts/graph"
	"contacts/graph/generated"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nicelogic/crdb"
	"github.com/nicelogic/pulsarclient"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/nicelogic/auth"
	"github.com/nicelogic/config"
	"github.com/nicelogic/errs"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	serviceConfig := contactsConfig.Config{
		Db_name:                 "contacts",
		Db_pool_connections_num: 4,
		Db_config_file_path:     "../../../test/config/config-crdb.yml",
		Pulsar_config_file_path: "../../../test/config/config-pulsar.yml",
		Path:                    "/",
		Listen_address:          ":80"}
	config.Init(constant.ConfigPath, &serviceConfig)
	crdbClient := crdb.Client{}
	err := crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Fatalf("crdb init err: %v", err)
	}
	pulsarClient := pulsarclient.Client{}
	err = pulsarClient.Init(serviceConfig.Pulsar_config_file_path,
		constant.PulsarTopic)
	if err != nil {
		log.Fatalf("Could not init Pulsar client: %v", err)
	}
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{CrdbClient: &crdbClient, PulsarClient: &pulsarClient}}))
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
		case graph.ContactsAddedMe == e.Error():
			fmt.Printf("ContactsAlreadyAddedU")
		default:
			err.Message = errs.ServerInternalError
		}
		return err
	})

	server.AddTransport(transport.POST{})
	server.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	server.Use(extension.Introspection{})
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle(serviceConfig.Path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", corsHandler.Handler(server))
	log.Printf("connect to http://%s%s for GraphQL playground", serviceConfig.Listen_address, serviceConfig.Path)
	log.Fatal(http.ListenAndServe(serviceConfig.Listen_address, router))
}
