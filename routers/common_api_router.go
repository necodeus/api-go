package routers

import (
	"github.com/gorilla/mux"
	"necodeo.com/m/v2/controllers/common_api"
)

func CommonApiRoutes(r *mux.Router) {
	domains := []string{"common-api.necodeo.com", "common-api.localhost"}

	for _, domain := range domains {
		subRouter := r.Host(domain).Subrouter()
		subRouter.HandleFunc("/", common_api.IndexHandler).Methods("GET")
	}
}
