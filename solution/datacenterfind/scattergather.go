package main

import (
	"log/slog"
	"sync"
)

func ScatterGather(resourceName string, finders []Finder) []Result {
	ch := make(chan Result, len(finders))
	var wg sync.WaitGroup

	// Scatter
	for _, finder := range finders {
		wg.Go(func() {
			slog.Info("starting find", slog.Any("datacenter", finder))
			ch <- Result{
				datacenter: finder,
				found:      finder.Find(resourceName),
			}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Gather
	var results []Result
	for result := range ch {
		results = append(results, result)
	}
	return results
}
