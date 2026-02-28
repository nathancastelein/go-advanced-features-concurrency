package main

import (
	"context"
	"log/slog"
	"time"
)

func Hedged(resourceName string, finders []Finder) {
	results := make(chan bool)
	defer close(results)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, finder := range finders {
		timer := time.NewTimer(75 * time.Millisecond)
		defer timer.Stop()
		go func() {
			slog.Info("launching find", slog.Any("datacenter", finder))
			found, err := finder.FindWithContext(ctx, resourceName)
			if err == nil {
				results <- found
			}
		}()

		select {
		case result := <-results:
			slog.Info("got result", slog.Bool("found", result))
			return
		case <-timer.C:
			continue
		}
	}
}
