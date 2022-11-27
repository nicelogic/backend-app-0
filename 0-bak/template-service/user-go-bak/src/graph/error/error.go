package error

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/nicelogic/errs"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	UserNotExist = "user not exist"
)

func HandleError(server *handler.Server){
	server.SetRecoverFunc(func(ctx context.Context, panicErr interface{}) error {
		log.Printf("panic: %v\n", panicErr)
		err := &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    errs.ServerInternalError,
			Extensions: map[string]interface{}{},
		}
		return err
	})
	server.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		log.Printf("error: %v\n", e)
		err := graphql.DefaultErrorPresenter(ctx, e)
		switch {
		case err.Message == UserNotExist:
			log.Printf(err.Message)
		default:
			err.Message = errs.ServerInternalError
		}
		return err
	})	
}