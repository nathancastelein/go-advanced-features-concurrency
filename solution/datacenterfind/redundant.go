package main

import (
	"context"
	"log/slog"
)

func Redundant(resourceName string, finders []Finder) []Result {
	ch := make(chan Result, len(finders))
	ctx, cancel := context.WithCancel(context.Background())

	for _, finder := range finders {
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
	}

	result := <-ch
	cancel()
	close(ch)
	return []Result{result}
}
