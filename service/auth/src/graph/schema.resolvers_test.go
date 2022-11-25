package graph_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestJwt(t *testing.T) {
	mySigningKey := []byte("AllYourBase")
	type User struct {
		id string
	}
	type MyCustomClaims struct {
		user *User
		jwt.RegisteredClaims
	}

	user := &User{ id: "test"}
	// Create the claims
	claims := MyCustomClaims{
		user, 
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "luojm",
			Subject:   "auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

}
