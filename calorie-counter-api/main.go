package main

import (
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"log"
)

func init() {
	if err := db.Migrate("gouser", "gouser", "127.0.0.1", "calories_dev"); err != nil {
		log.Fatal("[db-migrate] Error while applying migrations: ", err.Error())
	}
}

func main() {
	fmt.Println("Hello")
}
