package routers

import (
	"github.com/gorilla/mux"
	"necodeo.com/m/v2/rest/controllers/images"
)

func ImagesRoutes(r *mux.Router) {
	domains := []string{"images.necodeo.com", "images.localhost"}

	for _, domain := range domains {
		subRouter := r.Host(domain).Subrouter()
		subRouter.HandleFunc("/", images.IndexHandler).Methods("GET")
	}
}
