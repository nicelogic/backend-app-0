package main

import (
	"log"
	"net/http"
	userConfig "user/config"
	"user/constant"
	"user/graph"
	"user/graph/dependence"
	usererror "user/graph/error"
	"user/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/nicelogic/config"
	"github.com/rs/cors"
)

func main() {
	serviceConfig := userConfig.Config{}
	config.Init(constant.ConfigPath, &serviceConfig)
	authUtil, crdbClient, err := dependence.Init(&serviceConfig)
	if err != nil {
		log.Printf("dependence init err: %v\n", err)
	}
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					AuthUtil:   authUtil,
					CrdbClient: crdbClient,
				}}))
	usererror.HandleError(server)
	server.AddTransport(transport.POST{})
	server.Use(extension.Introspection{})
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})
	router := chi.NewRouter()
	router.Use(authUtil.Middleware())
	router.Handle(serviceConfig.Path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", corsHandler.Handler(server))
	log.Printf("connect to http://%s%s for GraphQL playground", serviceConfig.Listen_address, serviceConfig.Path)
	log.Fatal(http.ListenAndServe(serviceConfig.Listen_address, router))
}
