package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DeepAung/apiplustech-training/pokedex/database"
	"github.com/DeepAung/apiplustech-training/pokedex/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var (
	InvalidIntError       = errors.New("invalid integer")
	InvalidJsonFieldError = func(field string) error {
		return errors.New(fmt.Sprintf("invalid json field %q", field))
	}
)

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
		return nil, InvalidJsonFieldError("types")
	}

	var abilities []string
	if err := json.Unmarshal([]byte(pokemon.Abilities), &abilities); err != nil {
		return nil, InvalidJsonFieldError("abilities")
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
