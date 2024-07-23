package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	configs "github.com/F1sssss/go_ecom/config"
)

func NewDB() (*sql.DB, error) {

	DSN := createDSN(
		configs.Envs.DBUser,
		configs.Envs.DBPassword,
		configs.Envs.DBHost,
		configs.Envs.DBPort,
		configs.Envs.DBName,
	)

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	log.Println("Connected to database")

	return db, nil

}

func createDSN(user, password, host, port, dbname string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
}
