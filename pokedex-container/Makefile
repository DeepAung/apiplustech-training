air:
	air -c .air.toml

db.create:
	docker create --name pokedex-postgresql --env-file=.env.dev -p 5432:5432 bitnami/postgresql:latest
db.start:
	docker start -i pokedex-postgresql
db.migrate:
	docker cp ./database/schema.sql pokedex-postgresql:/docker-entrypoint-initdb.d/schema.sql
	docker exec -i pokedex-postgresql psql -U myuser -d mydb -f /docker-entrypoint-initdb.d/schema.sql

graphql.generate:
	go run github.com/99designs/gqlgen generate
