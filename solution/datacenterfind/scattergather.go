package main

import (
	"log/slog"
	"sync"
)

func ScatterGather(resourceName string, finders []Finder) {
	results := make(chan Result, len(finders))
	var wg sync.WaitGroup

	// Scatter
	for _, finder := range finders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			slog.Info("starting find", slog.Any("datacenter", finder))
			results <- Result{
				datacenter: finder,
				found:      finder.Find(resourceName),
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Gather
	for result := range results {
		slog.Info("got result", slog.Any("datacenter", result.datacenter), slog.Bool("found", result.found))
	}
}
