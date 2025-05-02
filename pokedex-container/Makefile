air:
	air -c .air.toml

db.generate:
	sqlc generate
db.migrate:
	sqlite3 pokedex.db < database/schema.sql

graphql.generate:
	go run github.com/99designs/gqlgen generate
