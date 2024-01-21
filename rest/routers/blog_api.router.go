package routers

import (
	"github.com/gorilla/mux"
	"necodeo.com/m/v2/rest/controllers/api/blog"
)

func BlogApiRoutes(r *mux.Router) {
	domains := []string{"blog-api.necodeo.com", "blog-api.localhost"}

	for _, domain := range domains {
		subRouter := r.Host(domain).Subrouter()
		subRouter.HandleFunc("/", blog.IndexHandler).Methods("GET")
	}
}
