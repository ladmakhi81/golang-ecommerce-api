package storage

import (
	"database/sql"
	"fmt"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func buildConnectionString(dbConfig config.MainConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.DatabaseUsername,
		dbConfig.DatabasePassword,
		dbConfig.DatabaseHost,
		dbConfig.DatabasePort,
		dbConfig.DatabaseName,
	)
}

func NewStorage(config config.MainConfig) *Storage {
	connection := buildConnectionString(config)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic("database is not connected")
	}
	if pingErr := db.Ping(); pingErr != nil {
		panic("database is not connected")
	}
	return &Storage{DB: db}
}
