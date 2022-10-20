package main

import (
	"context"
	"log"
	"net/http"
	"user/config"
	"user/graph"
	"user/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/nicelogic/auth"
	"github.com/nicelogic/configer"
	"github.com/nicelogic/constant/common_error"
	"github.com/vektah/gqlparser/v2/gqlerror"
)



func main() {

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	server.SetRecoverFunc(func(ctx context.Context, panicErr interface{}) error {
		log.Print(panicErr)
		err := &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    common_error.InternalServerError,
			Extensions: map[string]interface{}{},
		}
		return err
	})

	userConfig := config.Config{Path: "/"}
	configer.Init("/etc/app-0/config-user/config-user.yml", &userConfig)
	path := userConfig.Path
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle(path, playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)
	log.Printf("connect to http://localhost" + path + " for GraphQL playground")

	log.Fatal(http.ListenAndServe(":80", router))
}
