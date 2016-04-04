package db

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"bytes"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"github.com/tanel/dbmigrate"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/net/context"
	"gopkg.in/pg.v3"
)

var db *pg.DB

type Config struct {
	User     string
	Password string
	Host     string
	Database string
}

// Init starts the connections with the database
func Init(ctx context.Context, cfg Config) *pg.DB {
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
func Migrate(ctx context.Context, cfg Config) error {
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

// NewUUID generates a valid uuid v4
func NewUUID() string {
	return uuid.NewV4().String()
}

// IsUUID checks validity of UUID
func IsUUID(s string) bool {
	if len(s) != 36 {
		return false
	}
	if _, err := uuid.FromString(s); err != nil {
		return false
	}
	return true
}

// Hash uses scrypt to Hash a string
func Hash(password string, salt []byte) (string, error) {
	buffer := bytes.NewBuffer(nil)
	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err == nil {
		fmt.Fprintf(buffer, "%x", key)
	}
	return buffer.String(), err
}
