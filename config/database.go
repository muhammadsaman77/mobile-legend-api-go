package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	err:= godotenv.Load(".env")
	if err!= nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")
  dbName := os.Getenv("DB_NAME")
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	
	db,err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	return db

}