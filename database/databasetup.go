package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB_HOST, DB_NAME, DB_PASSWORD, DB_USER string
var DB_PORT int
var ErrDataBaseConnection error
var ErrDataBasePing error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	ErrDataBaseConnection = errors.New("can not connect to database")
}

func NewDatabasePool() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	db, err := sqlx.ConnectContext(ctx, "postgres", psqlconn)
	defer cancel()
	if err != nil {
		return nil
	}
	log.Println("Database successfully connection")
	return db
}
