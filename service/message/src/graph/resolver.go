package graph

import (
	contactsConfig "contacts/config"

	"github.com/nicelogic/crdb"
	"github.com/nicelogic/pulsarclient"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Config *contactsConfig.Config
	CrdbClient *crdb.Client
	PulsarClient *pulsarclient.Client
}
