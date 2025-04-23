-- name: ListPokemons :many
SELECT * FROM pokemons;

-- name: GetPokemonByID :one
SELECT * FROM pokemons WHERE id = ? LIMIT 1;

-- name: GetPokemonByName :one
SELECT * FROM pokemons WHERE name = ? LIMIT 1;

-- name: CreatePokemon :one
INSERT INTO pokemons (name, description, category, types, abilities)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdatePokemon :one
UPDATE pokemons SET
name = ?, description = ?, category = ?, types = ?, abilities = ?
WHERE id = ?
RETURNING *;


-- name: DeletePokemon :exec
DELETE FROM pokemons WHERE id = ?;
