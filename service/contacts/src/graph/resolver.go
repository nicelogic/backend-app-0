package graph

import (
	"crdb"

	"github.com/apache/pulsar-client-go/pulsar"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CrdbClient *crdb.Client
	PulsarClient *pulsar.Client
}
