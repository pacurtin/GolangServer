package services

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pacurtin/GolangServer/settings"
)

func InitDB() () {

	db, err := sql.Open("mysql", settings.Get().DataSourceName)

	if err != nil {
		log.Panic(err)
	}

	rows, err := db.Query("SELECT password from users where username = 'paddy'")

	if err != nil {
		log.Panic(err)
	}

	print(rows.Next())

	defer db.Close()

}
