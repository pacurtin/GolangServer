package routers

import (
	"github.com/pacurtin/GolangServer/controllers"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")

	return router
}
