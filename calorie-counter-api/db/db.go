package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tanel/dbmigrate"
	"gopkg.in/pg.v3"
	"log"
	"path/filepath"
	"time"
)

var db *pg.DB

type Config struct {
	User     string
	Password string
	Host     string
	Database string
}

// Init starts the connections with the database
func Init(cfg Config) *pg.DB {
	if db == nil {
		log.Println("[init-db]", "user", cfg.User, "host", cfg.Host, "database", cfg.Database)
		db = pg.Connect(&pg.Options{
			User:               cfg.User,
			Password:           cfg.Password,
			Host:               cfg.Host,
			Database:           cfg.Database,
			WriteTimeout:       3 * time.Second,
			ReadTimeout:        30 * time.Second,
			IdleTimeout:        30 * time.Second,
			IdleCheckFrequency: 30 * time.Second,
		})
	}
	return db
}

// Migrate applies migrations in db/migrations directory
func Migrate(cfg Config) error {
	log.Println("[db-migrate]", "user", cfg.User, "host", cfg.Host, "database", cfg.Database)
	pgDb, err := sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
			cfg.User, cfg.Password, cfg.Host, cfg.Database))
	if err != nil {
		log.Println("ERROR - unable to migrate database.", "Error", err.Error())
		return err
	}
	path := filepath.Join("db", "migrations")
	log.Println("Path", path)
	if err := dbmigrate.Run(pgDb, path); err != nil {
		return err
	}
	return nil
}
