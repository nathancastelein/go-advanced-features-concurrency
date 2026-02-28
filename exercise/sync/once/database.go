package main

import (
	"log/slog"

	"github.com/google/uuid"
)

type DB struct{ id string }

var (
	databaseConnection *DB
)

func GetDB() *DB {
	if databaseConnection == nil {
		slog.Info("connecting to database")
		databaseConnection = &DB{id: uuid.New().String()}
	}

	return databaseConnection
}

func (d DB) LogValuer() slog.Value {
	return slog.StringValue(d.id)
}
