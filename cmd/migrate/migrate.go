package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // register the postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // register the file source
	"github.com/joho/godotenv"
)

func main() {
	// 0. Load Environment Variables
	if err := godotenv.Load("config/local.env"); err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}

	// 1. Database Connection String
	host := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "user")
	dbPass := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "dexter_db")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		dbUser,
		dbPass,
		host,
		dbPort,
		dbName,
	)

	m, err := migrate.New(
		"file://./db/migration",
		connStr,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	// 2. Perform Migration
	if len(os.Args) > 1 && os.Args[1] == "down" {
		log.Println("Starting downward migration...")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration down failed: %v", err)
		}
		log.Println("Migration down finished")
	} else {
		log.Println("Starting upward migration...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration up failed: %v", err)
		}
		log.Println("Migration up finished")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
