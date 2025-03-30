package services

import (
	"context"
	"gocloud/internals/models"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func GetDb() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	createtable(conn)
	return conn
}

func createtable(conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    )
`)
	if err != nil {
		log.Fatalf("Unable to create table: %v", err)
	}

}

func InsertUser(conn *pgx.Conn, user models.User) error {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)

	return err
}
func GetUserByEmail(conn *pgx.Conn, email string) (*models.User, error) {
	var user models.User
	err := conn.QueryRow(context.Background(), "SELECT  name, email, password FROM users WHERE email=$1", email).
		Scan(&user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
