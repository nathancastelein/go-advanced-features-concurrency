package main

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

func ErrGroup(resourceName string, finders []Finder) {
	results := make(chan Result, len(finders))
	errGroup, ctx := errgroup.WithContext(context.Background())

	for _, finder := range finders {
		errGroup.Go(func() error {
			found, err := finder.FindWithError(ctx, resourceName)
			if err != nil {
				return err
			}
			results <- Result{
				datacenter: finder,
				found:      found,
			}
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		slog.Error("an error occured", slog.String("error", err.Error()))
		return
	}

	close(results)
	for result := range results {
		slog.Info("got result", slog.Any("datacenter", result.datacenter), slog.Bool("found", result.found))
	}
}
