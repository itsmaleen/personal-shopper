package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/itsmaleen/personal-shopper/backend/graph"
	"github.com/itsmaleen/personal-shopper/backend/graph/generated"
	"github.com/joho/godotenv"

	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:" + defaultPort},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		Debug:            true,
	})

	Database := graph.Connect()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "localhost"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	corsHandler := cors.Default().Handler(mux)
	corsHandler = c.Handler(corsHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
