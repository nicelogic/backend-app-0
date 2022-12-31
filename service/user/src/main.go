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
	authUtil, crdbClient, minioClient, err := dependence.Init(&serviceConfig)
	if err != nil {
		log.Printf("dependence init err: %v\n", err)
	}
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Config: &serviceConfig,
					AuthUtil:   authUtil,
					CrdbClient: crdbClient,
					MinioClient: minioClient,
				}}))
	usererror.HandleError(server)
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
	router.Use(authUtil.Middleware())
	router.Handle(serviceConfig.Path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)
	log.Printf("connect to http://%s%s for GraphQL playground", serviceConfig.Listen_address, serviceConfig.Path)
	log.Fatal(http.ListenAndServe(serviceConfig.Listen_address, router))
}
