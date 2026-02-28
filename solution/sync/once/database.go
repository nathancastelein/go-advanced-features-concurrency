package main

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

type DB struct{ id string }

var (
	databaseConnection *DB
	once               sync.Once
)

func GetDB() *DB {
	once.Do(func() {
		slog.Info("connecting to database")
		databaseConnection = &DB{id: uuid.New().String()}
	})

	return databaseConnection
}

func (d DB) LogValuer() slog.Value {
	return slog.StringValue(d.id)
}
