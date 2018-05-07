package main

import (
	"github.com/pacurtin/GolangServer/routers"
	"github.com/pacurtin/GolangServer/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
