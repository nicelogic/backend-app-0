package main

import (
	authConfig "auth/config"
	"auth/constant"
	"auth/graph"
	"auth/graph/dependence"
	autherror "auth/graph/error"
	"auth/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/nicelogic/config"
	"github.com/rs/cors"
)

func main() {
	serviceConfig := authConfig.Config{}
	config.Init(constant.ConfigPath, &serviceConfig)
	authUtil, crdbClient, err := dependence.Init(&serviceConfig)
	if err != nil {
		log.Printf("dependence init err: %v\n", err)
	}
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Config:       &serviceConfig,
					AuthUtil: authUtil,
					CrdbClient:   crdbClient,
					}}))
	autherror.HandleError(server)
	server.AddTransport(transport.POST{})
	server.Use(extension.Introspection{})
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Handle(serviceConfig.Path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)
	log.Printf("connect to http://%s%s for GraphQL playground", serviceConfig.Listen_address, serviceConfig.Path)
	log.Fatal(http.ListenAndServe(serviceConfig.Listen_address, router))
}
