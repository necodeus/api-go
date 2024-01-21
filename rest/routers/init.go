package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	ImagesRoutes(r)
	BlogApiRoutes(r)
	GraphRoutes(r)

	return r
}
