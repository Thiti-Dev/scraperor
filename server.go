package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Thiti-Dev/scraperor/core"
	"github.com/Thiti-Dev/scraperor/graph"
	"github.com/Thiti-Dev/scraperor/graph/generated"
	"github.com/Thiti-Dev/scraperor/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	// ─── LOADING ENV FILE ON LOCAL DEVELOPMENT ──────────────────────────────────────
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("[DEBUG]: No .env found")
	}
	// ────────────────────────────────────────────────────────────────────────────────

	DB := postgres.New(&pg.Options{
		Addr:    os.Getenv("ELEPHANTSQL_ADDRESS"),
		User:  	os.Getenv("ELEPHANTSQL_USER"),
		Password: os.Getenv("ELEPHANTSQL_PASSWORD"),
		Database: os.Getenv("ELEPHANTSQL_DATABASE"),
		
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UsersRepository: postgres.UsersRepository{DB:DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.HandleFunc("/scrape", core.ScrapeHandle())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
