-- name: ListPokemons :many
SELECT * FROM pokemons;

-- name: GetPokemonByID :one
SELECT * FROM pokemons WHERE id = $1 LIMIT 1;

-- name: GetPokemonByName :one
SELECT * FROM pokemons WHERE name = $1 LIMIT 1;

-- name: CreatePokemon :one
INSERT INTO pokemons (name, description, category, types, abilities)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdatePokemon :one
UPDATE pokemons SET
name = $2, description = $3, category = $4, types = $5, abilities = $6
WHERE id = $1
RETURNING *;


-- name: DeletePokemon :exec
DELETE FROM pokemons WHERE id = $1;
