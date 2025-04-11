package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DeepAung/apiplustech-training/review-it/graph"
	"github.com/DeepAung/apiplustech-training/review-it/graph/model"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			DB: graph.Database{
				MoviesTable:  make(map[string]model.Movie),
				ReviewsTable: make(map[string]model.Review),
			},
		}}),
	)

	r.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
