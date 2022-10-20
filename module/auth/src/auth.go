package auth

import (
	"context"
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

type User struct {
	Id string
}

func GetUser(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

func userFromJwt(tokenString string) (user *User, err error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		content, err := os.ReadFile("/etc/app-0/secret-jwt/jwt-publickey")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		jwtPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(content)
		if err != nil{
			fmt.Println(err)
			return nil, err
		}
		return jwtPublicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		_, err = fmt.Println(err)
		return
	}
	userMap, ok := claims["user"].(map[string]interface{})
	if !ok {
		_, err = fmt.Printf("claims[user] is not map[string]interface{}\n")
		return
	}
	id, ok := userMap["id"].(string)
	if !ok {
		_, err = fmt.Printf("userMap[id] is not string\n")
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
			if len(splitToken) == 2 {
				jwtToken := splitToken[1]
				user, err := userFromJwt(jwtToken)
				if err == nil {
					ctx := context.WithValue(request.Context(), userCtxKey, user)
					request = request.WithContext(ctx)
				}
			}
			next.ServeHTTP(writer, request)
		})
	}
}