package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
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
	repo *Repository
}

type Repository struct {
	queries *database.Queries
}

func NewResolver(queries *database.Queries) *Resolver {
	return &Resolver{
		repo: &Repository{
			queries: queries,
		},
	}
}

func (r *Repository) CreatePokemon(
	ctx context.Context,
	input model.PokemonInput,
) (*model.Pokemon, error) {
	types, err := json.Marshal(&input.Types)
	if err != nil {
		return nil, InvalidJsonFieldError("types")
	}

	abilities, err := json.Marshal(&input.Abilities)
	if err != nil {
		return nil, InvalidJsonFieldError("abilities")
	}

	result, err := r.queries.CreatePokemon(ctx, database.CreatePokemonParams{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Types:       string(types),
		Abilities:   string(abilities),
	})
	if err != nil {
		return nil, err
	}

	return r.convertPokemonDBToGraphql(&result)
}

func (r *Repository) UpdatePokemon(
	ctx context.Context,
	id int64,
	input model.PokemonInput,
) (*model.Pokemon, error) {
	types, err := json.Marshal(&input.Types)
	if err != nil {
		return nil, InvalidJsonFieldError("types")
	}

	abilities, err := json.Marshal(&input.Abilities)
	if err != nil {
		return nil, InvalidJsonFieldError("abilities")
	}

	result, err := r.queries.UpdatePokemon(ctx, database.UpdatePokemonParams{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Types:       string(types),
		Abilities:   string(abilities),
		ID:          id,
	})
	if err != nil {
		return nil, err
	}

	return r.convertPokemonDBToGraphql(&result)
}

func (r *Repository) DeletePokemon(ctx context.Context, id int64) (bool, error) {
	err := r.queries.DeletePokemon(ctx, id)
	return err == nil, err
}

func (r *Repository) ListPokemons(ctx context.Context) ([]*model.Pokemon, error) {
	result, err := r.queries.ListPokemons(ctx)
	if err != nil {
		return nil, err
	}

	pokemons := make([]*model.Pokemon, len(result))
	for i, item := range result {
		pokemons[i], err = r.convertPokemonDBToGraphql(&item)
		if err != nil {
			return nil, err
		}
	}

	return pokemons, nil
}

func (r *Repository) GetPokemonByID(ctx context.Context, id int64) (*model.Pokemon, error) {
	result, err := r.queries.GetPokemonByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return r.convertPokemonDBToGraphql(&result)
}

func (r *Repository) GetPokemonByName(ctx context.Context, name string) (*model.Pokemon, error) {
	result, err := r.queries.GetPokemonByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return r.convertPokemonDBToGraphql(&result)
}

// -------------------------------------------------------------------------- //

func (r *Repository) convertPokemonDBToGraphql(pokemon *database.Pokemon) (*model.Pokemon, error) {
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
