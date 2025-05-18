package database

import (
	"fmt"

	"github.com/DeepAung/apiplustech-training/pokedex/graph/model"
	"github.com/uptrace/bun"
)

type Pokemon struct {
	bun.BaseModel `bun:"table:pokemons"`
	Id            int64                  `bun:"id,pk,autoincrement"`
	Name          string                 `bun:"name,notnull,unique"`
	Description   string                 `bun:"description,notnull"`
	Category      string                 `bun:"category,notnull"`
	Types         []model.PokemonType    `bun:"types,array,notnull"`
	Abilities     []model.PokemonAbility `bun:"abilities,array,notnull"`
}

func (p *Pokemon) ToGraphql() *model.Pokemon {
	return &model.Pokemon{
		ID:          fmt.Sprint(p.Id),
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Types:       p.Types,
		Abilities:   p.Abilities,
	}
}

func PokemonToDbModel(p *model.Pokemon) *Pokemon {
	return &Pokemon{
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Types:       p.Types,
		Abilities:   p.Abilities,
	}
}

func PokemonInputToDbModel(p *model.PokemonInput) *Pokemon {
	return &Pokemon{
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Types:       p.Types,
		Abilities:   p.Abilities,
	}
}
