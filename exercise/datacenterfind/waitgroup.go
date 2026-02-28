package main

import (
	"log/slog"
)

func WaitGroup(resourceName string, finders []Finder) {
	for _, finder := range finders {
		slog.Info("starting find", slog.Any("datacenter", finder))
		found := finder.Find(resourceName)
		slog.Info("got result", slog.Any("datacenter", finder), slog.Bool("found", found))
	}
}
