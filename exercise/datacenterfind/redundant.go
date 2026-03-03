package main

import (
	"context"
	"log/slog"
	"sync"
)

func Redundant(resourceName string, finders []Finder) []Result {
	ch := make(chan Result, len(finders))
	var wg sync.WaitGroup

	for _, finder := range finders {
		wg.Go(func() {
			slog.Info("starting find", slog.Any("datacenter", finder))
			found, err := finder.FindWithContext(context.TODO(), resourceName)
			if err == nil {
				ch <- Result{
					datacenter: finder,
					found:      found,
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var results []Result
	for result := range ch {
		results = append(results, result)
	}
	return results
}
