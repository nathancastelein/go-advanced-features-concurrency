package main

import (
	"context"
	"log/slog"
	"time"
)

func Hedged(resourceName string, finders []Finder) []Result {
	ch := make(chan Result)
	defer close(ch)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, finder := range finders {
		timer := time.NewTimer(75 * time.Millisecond)
		defer timer.Stop()
		go func() {
			slog.Info("launching find", slog.Any("datacenter", finder))
			found, err := finder.FindWithContext(ctx, resourceName)
			if err == nil {
				ch <- Result{
					datacenter: finder,
					found:      found,
				}
			}
		}()

		select {
		case result := <-ch:
			return []Result{result}
		case <-timer.C:
			continue
		}
	}
	return nil
}
