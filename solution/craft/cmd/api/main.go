package main

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/nathancastelein/go-advanced-features-concurrency/solution/craft/pkg/http"
	"github.com/nathancastelein/go-advanced-features-concurrency/solution/craft/pkg/storage"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://localhost:5432/user")
	if err != nil {
		slog.Error("error while connecting to database", slog.Any("error", err))
		return
	}

	server := http.NewServer(8080, storage.NewUserSQL(db))
	if err := server.Start(); err != nil {
		slog.Error("error while starting server", slog.Any("error", err))
		return
	}
}
