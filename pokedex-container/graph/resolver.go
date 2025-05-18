package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DeepAung/apiplustech-training/pokedex/database"
	"github.com/DeepAung/apiplustech-training/pokedex/graph/model"
	"github.com/uptrace/bun"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var (
	InvalidIntError             = errors.New("invalid integer")
	NotUniqueTypeOrAbilityError = errors.New("not unique type or ability")
	UniqueNameError             = errors.New("name already exists")
	IdNotFoundError             = errors.New("pokemon id not found")
	NameNotFoundError           = errors.New("pokemon name not found")
)

type Resolver struct {
	repo *Repository
}

type Repository struct {
	db *bun.DB
}

func NewResolver(db *bun.DB) *Resolver {
	return &Resolver{
		repo: &Repository{
			db: db,
		},
	}
}

func (r *Repository) CreatePokemon(
	ctx context.Context,
	input model.PokemonInput,
) (*model.Pokemon, error) {
	var result database.Pokemon
	model := database.PokemonInputToDbModel(&input)

	if err := r.db.NewInsert().Model(model).Returning("*").Scan(ctx, &result); err != nil {
		if err.Error() == "ERROR: new row for relation \"pokemons\" violates check constraint \"pokemons_check\" (SQLSTATE=23514)" {
			return nil, NotUniqueTypeOrAbilityError
		}
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"pokemons_name_key\" (SQLSTATE=23505)" {
			return nil, UniqueNameError
		}
		return nil, err
	}

	return result.ToGraphql(), nil
}

func (r *Repository) UpdatePokemon(
	ctx context.Context,
	id int32,
	input model.PokemonInput,
) (*model.Pokemon, error) {
	var result database.Pokemon
	model := database.PokemonInputToDbModel(&input)

	if err := r.db.NewUpdate().Model(model).Where("id = ?", id).Returning("*").Scan(ctx, &result); err != nil {
		return nil, err
	}

	return result.ToGraphql(), nil
}

func (r *Repository) DeletePokemon(ctx context.Context, id int32) (bool, error) {
	result, err := r.db.NewDelete().Model((*database.Pokemon)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return false, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if n == 0 {
		return false, IdNotFoundError
	}
	return true, nil
}

func (r *Repository) ListPokemons(ctx context.Context) ([]*model.Pokemon, error) {
	var result []database.Pokemon
	err := r.db.NewSelect().Model((*database.Pokemon)(nil)).Scan(ctx, &result)
	if err != nil {
		return nil, err
	}

	pokemons := make([]*model.Pokemon, len(result))
	for i := range len(result) {
		pokemons[i] = result[i].ToGraphql()
	}
	return pokemons, nil
}

func (r *Repository) GetPokemonByID(ctx context.Context, id int32) (*model.Pokemon, error) {
	var result database.Pokemon
	err := r.db.NewSelect().
		Model((*database.Pokemon)(nil)).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx, &result)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, IdNotFoundError
		}
		return nil, err
	}

	return result.ToGraphql(), nil
}

func (r *Repository) GetPokemonByName(ctx context.Context, name string) (*model.Pokemon, error) {
	var result database.Pokemon
	err := r.db.NewSelect().
		Model((*database.Pokemon)(nil)).
		Where("name = ?", name).
		Limit(1).
		Scan(ctx, &result)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, NameNotFoundError
		}
		return nil, err
	}

	return result.ToGraphql(), nil
}
