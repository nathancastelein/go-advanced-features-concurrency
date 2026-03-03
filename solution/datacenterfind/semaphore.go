package main

import (
	"log/slog"
	"sync"
)

func Semaphore(resourceName string, finders []Finder) []Result {
	results := make([]Result, len(finders))
	var wg sync.WaitGroup
	maxParallelAccesses := 2
	semaphore := make(chan struct{}, maxParallelAccesses)
	defer close(semaphore)

	for i, finder := range finders {
		wg.Go(func() {
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

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
