package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"encoding/json"
	"fmt"

	"github.com/DeepAung/apiplustech-training/pokedex/database"
	"github.com/DeepAung/apiplustech-training/pokedex/graph/model"
)

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

func (r *Resolver) convertPokemonDBToGraphql(pokemon *database.Pokemon) (*model.Pokemon, error) {
	var types []model.PokemonType
	if err := json.Unmarshal([]byte(pokemon.Types), &types); err != nil {
		return nil, err
	}

	var abilities []string
	if err := json.Unmarshal([]byte(pokemon.Abilities), &abilities); err != nil {
		return nil, err
	}

	return &model.Pokemon{
		ID:          fmt.Sprint(pokemon.ID),
		Name:        pokemon.Name,
		Description: pokemon.Description,
		Category:    pokemon.Category,
		Types:       types,
		Abilities:   abilities,
	}, nil
}
