package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB(){
	if err:=godotenv.Load();err!=nil{
		log.Println("No .env file found, reading from system env")
	}

	connStr:=os.Getenv("DATABASE_URL")
	if connStr==""{
		log.Fatal("DATABASE_URL is not set in .env")
	}

	var err error
	DB, err= pgxpool.New(context.Background(),connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Connected to PostgreSQL!")

	createTableQuery:=`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
	_,err=DB.Exec(context.Background(),createTableQuery)
	if err!=nil{
		log.Fatalf("Failed to create users table: %v\n",err)
	}
	log.Println("Database tables initialized.")
}