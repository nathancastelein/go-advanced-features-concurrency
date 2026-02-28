package main

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

type fakeFinder struct {
	delay time.Duration
	calls int
}

func (f *fakeFinder) Find(resourceName string) bool {
	f.calls++
	time.Sleep(f.delay)
	return true
}

func (f *fakeFinder) FindWithContext(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (f *fakeFinder) FindWithError(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (f *fakeFinder) Weight() int64 { return 1 }

func TestWaitGroup(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		finders := []Finder{
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 200 * time.Millisecond},
			&fakeFinder{delay: 300 * time.Millisecond},
		}

		start := time.Now()
		WaitGroup("server-1", finders)
		elapsed := time.Since(start)

		// All finders must have been called
		for i, f := range finders {
			require.Equal(t, 1, f.(*fakeFinder).calls, "finder %d should have been called once", i)
		}

		// With concurrency, total time should be ~300ms (slowest), not 600ms (sum)
		require.Less(t, elapsed, 400*time.Millisecond, "finders should run concurrently, not sequentially")
	})
}
