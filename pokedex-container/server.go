package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DeepAung/apiplustech-training/pokedex/database"
	"github.com/DeepAung/apiplustech-training/pokedex/graph"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

var envPath = flag.String("env-file", "", "environment file path")

func main() {
	flag.Parse()
	cfg := loadConfig(*envPath)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.dbUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	queries := database.New(conn)
	resolver := graph.NewResolver(queries)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.appPort)
	if err := http.ListenAndServe(":"+cfg.appPort, r); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	dbUrl   string
	appPort string
}

func loadConfig(filePath string) config {
	if filePath != "" {
		_ = godotenv.Load(filePath)
	}

	return config{
		dbUrl:   os.Getenv("DB_URL"),
		appPort: os.Getenv("APP_PORT"),
	}
}
