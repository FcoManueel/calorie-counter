package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tanel/dbmigrate"
	"log"
	"path/filepath"
)

// Migrate applies migrations in db/migrate directory
func Migrate(user, password, host, database string) error {
	log.Println("[db-migrate]", "user", user, "host", host, "database", database)
	pgDb, err := sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
			user, password, host, database))
	if err != nil {
		log.Println("ERROR - unable to migrate database.", "Error", err.Error())
		return err
	}
	path := filepath.Join("db", "migrate")
	log.Println("Path", path)
	if err := dbmigrate.Run(pgDb, path); err != nil {
		return err
	}
	return nil
}
