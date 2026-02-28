package main

import (
	"context"
	"log/slog"
)

func Redundant(resourceName string, finders []Finder) {
	results := make(chan Result, len(finders))
	ctx, cancel := context.WithCancel(context.Background())

	for _, finder := range finders {
		go func() {
			slog.Info("launching find", slog.Any("datacenter", finder))
			found, err := finder.FindWithContext(ctx, resourceName)
			if err == nil {
				results <- Result{
					datacenter: finder,
					found:      found,
				}
			}
		}()
	}

	result := <-results
	cancel()
	close(results)
	slog.Info("got result", slog.Any("datacenter", result.datacenter), slog.Bool("found", result.found))
}
