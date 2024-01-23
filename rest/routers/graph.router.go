package routers

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"necodeo.com/m/v2/graph/posts"
)

func GraphRoutes(r *mux.Router) {
	domains := []string{"graph.necodeo.com", "graph.localhost"}

	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"posts": &graphql.Field{
				Type:    graphql.NewList(posts.PostType),
				Resolve: posts.PostsResolver,
			},
		},
	})

	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   false,
		GraphiQL: true,
	})

	for _, domain := range domains {
		subRouter := r.Host(domain).Subrouter()
		subRouter.Path("/").Handler(h)
	}
}
