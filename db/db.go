package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Database struct {
	DB *sql.DB
}

func InitDB() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(10)
	//DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database!")
	return &Database{DB: DB}, nil
}
