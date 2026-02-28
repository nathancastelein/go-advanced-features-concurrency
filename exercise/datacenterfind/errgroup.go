package main

import (
	"context"
	"log/slog"
	"sync"
)

func ErrGroup(resourceName string, finders []Finder) {
	results := make(chan Result, len(finders))
	var wg sync.WaitGroup

	for _, finder := range finders {
		wg.Add(1)

		go func() {
			found, err := finder.FindWithError(context.TODO(), resourceName)
			if err != nil {
				panic(err)
			}
			results <- Result{
				datacenter: finder,
				found:      found,
			}
		}()
	}

	wg.Wait()
	close(results)

	for result := range results {
		slog.Info("got result", slog.Any("datacenter", result.datacenter), slog.Bool("found", result.found))
	}
}
