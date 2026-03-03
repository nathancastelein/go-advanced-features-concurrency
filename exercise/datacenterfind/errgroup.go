package main

import (
	"context"
	"sync"
)

func ErrGroup(resourceName string, finders []Finder) ([]Result, error) {
	results := make(chan Result, len(finders))
	var wg sync.WaitGroup

	for _, finder := range finders {
		wg.Go(func() {
			found, err := finder.FindWithError(context.TODO(), resourceName)
			if err != nil {
				panic(err)
			}
			results <- Result{
				datacenter: finder,
				found:      found,
			}
		})
	}

	wg.Wait()
	close(results)

	var allResults []Result
	for result := range results {
		allResults = append(allResults, result)
	}
	return allResults, nil
}
