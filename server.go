package main

import (
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
)

const defaultPort = "8080"

func main() {

	DB := postgres.New(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "root",
		Database: "scraperor",
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
