package main

import (
	"context"
	"log/slog"
	"sync"

	"golang.org/x/sync/semaphore"
)

func WeightedSemaphore(resourceName string, finders []Finder) {
	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(30)
	ctx := context.Background()

	for _, finder := range finders {
		wg.Go(func() {
			if err := sem.Acquire(ctx, finder.Weight()); err != nil {
				slog.Error("fail to acquire semaphore", slog.String("error", err.Error()))
				return
			}
			defer sem.Release(finder.Weight())

			slog.Info("starting find", slog.Any("datacenter", finder))
			found := finder.Find(resourceName)
			slog.Info("got result", slog.Any("datacenter", finder), slog.Bool("found", found))
		})
	}

	wg.Wait()
}
