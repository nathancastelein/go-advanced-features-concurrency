package main

import (
	"flag"
	"log/slog"
	"time"
)

var action string

func init() {
	flag.StringVar(&action, "action", "help", "choose action to perform")
}

func main() {
	flag.Parse()
	resourceName := "server-1"
	finders := []Finder{SBG(), GRA(), BHS()}

	switch action {
	case "sequential":
		Sequential(resourceName, finders)
	case "waitgroup":
		results := WaitGroup(resourceName, finders)
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
	case "scattergather":
		results := ScatterGather(resourceName, finders)
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
	case "redundant":
		results := Redundant(resourceName, finders)
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
		// wait a bit to see the context cancellation
		time.Sleep(200 * time.Millisecond)
	case "errgroup":
		results, err := ErrGroup(resourceName, append(finders, DCError()))
		if err != nil {
			slog.Error("an error occured", slog.String("error", err.Error()))
		}
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
	case "hedged":
		results := Hedged(resourceName, finders)
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
	case "semaphore":
		results := Semaphore(resourceName, append(finders, RBX(), WAW()))
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
	case "wsemaphore":
		results := WeightedSemaphore(resourceName, append(finders, RBX(), WAW()))
		for _, r := range results {
			slog.Info("got result", slog.Any("datacenter", r.datacenter), slog.Bool("found", r.found))
		}
	default:
		slog.Error("unknown action", slog.String("action", action))
	}
}
