package error

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nicelogic/errs"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	UserExist = "user already exist"
	UserNotExist = "user not exist"
	PwdWrong = "password wrong"
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
		var jwtError *jwt.ValidationError
		hasJwtError := errors.As(e, &jwtError)
		err := graphql.DefaultErrorPresenter(ctx, e)
		switch {
		case hasJwtError && jwtError.Errors == jwt.ValidationErrorExpired:
			err.Message = errs.TokenExpired
		case hasJwtError:
			err.Message = errs.TokenInvalid
		case err.Message == UserExist:
		case err.Message == UserNotExist:
		case err.Message == PwdWrong:
			log.Printf(err.Message)
		default:
			err.Message = errs.ServerInternalError
		}
		return err
	})	
}