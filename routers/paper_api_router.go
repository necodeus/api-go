package routers

import (
	"github.com/gorilla/mux"
	"necodeo.com/m/v2/controllers/paper_api"
)

func PaperApiRoutes(r *mux.Router) {
	domains := []string{"paper-api.necodeo.com", "paper-api.localhost"}

	for _, domain := range domains {
		subRouter := r.Host(domain).Subrouter()
		subRouter.HandleFunc("/", paper_api.IndexHandler).Methods("GET")
	}
}
