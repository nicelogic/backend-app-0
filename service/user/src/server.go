package main

import (
	"log"
	"net/http"
	"user/config"
	"user/graph"
	"user/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nicelogic/configer"
)

func main() {
	userConfig := config.Config{Path: "/"}
	configer.Init("/etc/app-0/config-user/config-user.yml", &userConfig)
	path := userConfig.Path

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle(path, playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost" + path + " for GraphQL playground")
	log.Fatal(http.ListenAndServe(":80", nil))
}
