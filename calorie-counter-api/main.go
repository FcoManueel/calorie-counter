package main

import (
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/routes"
)

func init() {
	if err := db.Migrate("gouser", "gouser", "127.0.0.1", "calories_dev"); err != nil {
		log.Fatal("[db-migrate] Error while applying migrations: ", err.Error())
	}
}

func main() {
	root := routes.Init()
	http.ListenAndServe("localhost:8000", root)
}
