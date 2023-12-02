package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	CommonApiRoutes(r)
	PaperApiRoutes(r)
	return r
}
