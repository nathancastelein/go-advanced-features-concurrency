package main

import (
	"log/slog"
	"sync"
)

func ScatterGather(resourceName string, finders []Finder) {
	var wg sync.WaitGroup

	for _, finder := range finders {
		wg.Go(func() {
			slog.Info("starting find", slog.Any("datacenter", finder))
			found := finder.Find(resourceName)
			slog.Info("got result", slog.Any("datacenter", finder), slog.Bool("found", found))
		})
	}

	wg.Wait()
}
