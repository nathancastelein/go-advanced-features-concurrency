package main

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWaitGroup(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Arrange
		finders := []Finder{
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 200 * time.Millisecond},
			&fakeFinder{delay: 300 * time.Millisecond},
		}

		// Act
		start := time.Now()
		results := WaitGroup("server-1", finders)
		elapsed := time.Since(start)

		// Assert
		require.Len(t, results, 3, "should return one result per finder")
		for _, result := range results {
			require.True(t, result.found)
		}
		require.Less(t, elapsed, 400*time.Millisecond, "finders should run concurrently, not sequentially")
	})
}

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
