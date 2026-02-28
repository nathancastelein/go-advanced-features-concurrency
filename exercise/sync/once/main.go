package main

import (
	"log/slog"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db := GetDB()
			slog.Info("got database connection!", slog.Any("database", db))
		}()
	}

	wg.Wait()
}
