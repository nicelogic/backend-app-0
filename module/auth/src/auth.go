package auth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{name: "user"}
var errorCtxKey = &contextKey{name: "error"}

type User struct {
	Id string
}

func GetUser(ctx context.Context) (*User, error) {
	user, _ := ctx.Value(userCtxKey).(*User)
	err, _ := ctx.Value(errorCtxKey).(error)
	return user, err
}

var jwtPublicKey *rsa.PublicKey

func userFromJwt(tokenString string) (user *User, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("recovered error: %v", err)
		}
	}()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if jwtPublicKey == nil {
			sliceJwtPublicKey, err := os.ReadFile("/etc/app-0/secret-jwt/jwt-publickey")
			if err != nil {
				return nil, err
			}
			jwtPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(sliceJwtPublicKey)
			if err != nil {
				return nil, err
			}
		}

		return jwtPublicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return
	}
	userMap, ok := claims["user"].(map[string]interface{})
	if !ok {
		err = errors.New("claims[user] is not map[string]interface{}")
		return
	}
	id, ok := userMap["id"].(string)
	if !ok {
		err = errors.New("userMap[id] is not string")
		return
	}
	user = &User{Id: id}
	return
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			reqToken := request.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			var ctx context.Context
			if len(splitToken) != 2 {
				ctx = context.WithValue(request.Context(), errorCtxKey, errors.New("http header: Authorization's value invalid"))
			} else {
				jwtToken := splitToken[1]
				user, err := userFromJwt(jwtToken)
				if err != nil {
					fmt.Printf("userFromJwt: %v\n", err)
					ctx = context.WithValue(request.Context(), errorCtxKey, err)
				} else {
					ctx = context.WithValue(request.Context(), userCtxKey, user)
				}
			}
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}
