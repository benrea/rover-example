package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/benrea/rover-example/graphql/gqlgen"
	"github.com/benrea/rover-example/graphql/resolver"
	"github.com/go-chi/chi/v5"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	srv := handler.New(gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers: &resolver.Resolver{},
	}))
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	router := chi.NewRouter()
	router.Handle("/graphql", srv)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	addr := "0.0.0.0:8080"
	log.Info("serving on ", addr)
	if err := http.ListenAndServe(addr, c.Handler(router)); err != nil {
		log.Error("server terminated: ", err)
	}
}
