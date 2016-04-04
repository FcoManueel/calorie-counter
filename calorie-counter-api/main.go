package main

import (
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/routes"
	"golang.org/x/net/context"
)

func init() {
	dbConfig := db.Config{
		User:     "gouser",
		Password: "",
		Host:     "127.0.0.1",
		Database: "calories_dev",
	}
	ctx := context.Background()
	if err := db.Migrate(ctx, dbConfig); err != nil {
		log.Fatalln("[db-migrate] Error while applying migrations: ", err.Error())
	}
	db.Init(ctx, dbConfig)
}

func main() {
	root := routes.Init()
	http.ListenAndServe("localhost:8000", root)
}
