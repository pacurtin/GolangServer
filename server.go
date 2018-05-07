package main

import (
	"github.com/pacurtin/GolangServer/routers"
	"github.com/pacurtin/GolangServer/settings"
	"github.com/codegangsta/negroni"
	"net/http"
	"database/sql"
)

func main() {
	settings.Init()

	// TODO move this to its own service(?)
	db, err := sql.Open("mysql", "user:password@/dbname")

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
