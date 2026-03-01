package main

import (
	"log/slog"
	"sync"
)

func ScatterGather(resourceName string, finders []Finder) []Result {
	results := make([]Result, len(finders))
	var wg sync.WaitGroup

	for i, finder := range finders {
		wg.Go(func() {
			slog.Info("starting find", slog.Any("datacenter", finder))
			results[i] = Result{
				datacenter: finder,
				found:      finder.Find(resourceName),
			}
		})
	}

	wg.Wait()
	return results
}
