package graph

import "github.com/DeepAung/apiplustech-training/pokedex/database"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	queries *database.Queries
}

func NewResolver(queries *database.Queries) *Resolver {
	return &Resolver{
		queries: queries,
	}
}
