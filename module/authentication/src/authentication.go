package authentication

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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userMap := claims["user"].(map[string]interface{})
		fmt.Printf("claims[user]: %v\n", userMap)
		user = &User{Id: userMap["id"].(string)}
	} else {
		fmt.Println(err)
	}	

	return 
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			reqToken := request.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			if(len(splitToken) != 2){
				fmt.Println("invalid token: ", reqToken)	
			} else {
				jwtToken := splitToken[1]
				fmt.Println("token: ", jwtToken)
				user, err := userFromJwt(jwtToken)
				if err != nil{
					ctx := context.WithValue(request.Context(), userCtxKey, user)
					request = request.WithContext(ctx)
				}
			}
			next.ServeHTTP(writer, request)
		})
	}
}