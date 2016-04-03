package main

import (
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/routes"
)

func init() {
	dbConfig := db.Config{
		User:     "gouser",
		Password: "",
		Host:     "127.0.0.1",
		Database: "calories_dev",
	}

	if err := db.Migrate(dbConfig); err != nil {
		log.Fatalln("[db-migrate] Error while applying migrations: ", err.Error())
	}
	db.Init(dbConfig)
}

func main() {
	root := routes.Init()
	http.ListenAndServe("localhost:8000", root)
}
