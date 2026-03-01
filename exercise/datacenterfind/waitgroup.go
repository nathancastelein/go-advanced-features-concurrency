package main

import (
	"log/slog"
)

func WaitGroup(resourceName string, finders []Finder) []Result {
	var results []Result
	for _, finder := range finders {
		slog.Info("starting find", slog.Any("datacenter", finder))
		results = append(results, Result{
			datacenter: finder,
			found:      finder.Find(resourceName),
		})
	}
	return results
}
