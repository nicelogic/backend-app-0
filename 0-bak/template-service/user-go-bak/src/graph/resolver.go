package graph

import (
	"github.com/nicelogic/authutil"
	"github.com/nicelogic/crdb"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	AuthUtil *authutil.Auth
	CrdbClient *crdb.Client
}
