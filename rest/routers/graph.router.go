package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"necodeo.com/m/v2/graphql"
	"necodeo.com/m/v2/graphql/generated"
)

func GraphRoutes(r *mux.Router) {
	domains := []string{"graph.necodeo.com", "graph.localhost"}

	for _, domain := range domains {
		subRouter := r.Host(domain).Subrouter()

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

		subRouter.HandleFunc("/", playground.Handler("GraphQL playground", "/query")).Methods("GET")

		subRouter.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
			srv.ServeHTTP(w, r)
		})
	}
}
