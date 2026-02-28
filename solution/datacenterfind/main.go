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
		WaitGroup(resourceName, finders)
	case "scattergather":
		ScatterGather(resourceName, finders)
	case "redundant":
		Redundant(resourceName, finders)
		// wait a bit to see the context cancellation
		time.Sleep(200 * time.Millisecond)
	case "errgroup":
		ErrGroup(resourceName, append(finders, DCError()))
	case "hedged":
		Hedged(resourceName, finders)
	case "semaphore":
		Semaphore(resourceName, append(finders, RBX(), WAW()))
	case "wsemaphore":
		WeightedSemaphore(resourceName, append(finders, RBX(), WAW()))
	default:
		slog.Error("unknown action", slog.String("action", action))
	}
}
