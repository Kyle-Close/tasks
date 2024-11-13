package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("Error loading environment variables.")
	}

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	addr := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	// Capture connection properties.
	cfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Net:    "tcp",
		Addr:   addr,
		DBName: name,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, errors.New("Error connecting to the mysql database.")
	}

	pingErr := db.Ping()

	if pingErr != nil {
		return nil, errors.New("Error pinging the mysql database.")
	}

	fmt.Println("Connection to db successful!")

	return db, nil
}
