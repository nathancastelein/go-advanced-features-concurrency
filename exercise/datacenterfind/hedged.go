package main

import (
	"context"
	"log/slog"
)

func Hedged(resourceName string, finders []Finder) {
	results := make(chan bool, len(finders))
	ctx, cancel := context.WithCancel(context.Background())

	for _, finder := range finders {
		go func() {
			slog.Info("launching find", slog.Any("datacenter", finder))
			found, err := finder.FindWithContext(ctx, resourceName)
			if err == nil {
				results <- found
			}
		}()
	}

	found := <-results
	cancel()
	close(results)
	slog.Info("got result", slog.Bool("found", found))
}
