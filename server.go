package main

import (
	"github.com/pacurtin/GolangServer/routers"
	"github.com/pacurtin/GolangServer/settings"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/pacurtin/GolangServer/services"
)

func main() {
	settings.Init()
	services.InitDB()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
