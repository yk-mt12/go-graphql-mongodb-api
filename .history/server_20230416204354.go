package main

import (
	"net/http"
    "log"
    "os"

    "go-graphql-mongodb-api/database"
    "go-graphql-mongodb-api/graph"
    "go-graphql-mongodb-api/graph/generated"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/rs/cors"
    "github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	database.Connect("mongodb://localhost:27017/")
	c := cors.New(cors.Options{
		
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
