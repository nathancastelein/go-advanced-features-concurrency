package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func ErrGroup(resourceName string, finders []Finder) ([]Result, error) {
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
		return nil, err
	}

	close(results)
	var allResults []Result
	for result := range results {
		allResults = append(allResults, result)
	}
	return allResults, nil
}
