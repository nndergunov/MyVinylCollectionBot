package service

import (
	"database/sql"
	"fmt"

	db "github.com/nndergunov/tgBot/server/pkg/db/tgbot/public"
)

type DB struct {
	db *sql.DB
}

func NewDB(dbSource string) (*DB, error) {
	database, err := db.NewDB(dbSource)
	if err != nil {
		return nil, fmt.Errorf("database open: %w", err)
	}

	return &DB{
		db: database,
	}, nil
}