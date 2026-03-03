package main

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWeightedSemaphore(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Arrange: capacity 30, 5 finders with weight 10 each
		// Weighted semaphore (capacity 30, weight 10): 3 fit per batch -> 2 batches = 200ms
		// Fixed-slot semaphore (capacity 2): 2 fit per batch -> 3 batches = 300ms
		// Without semaphore (fully concurrent): 100ms
		finders := []Finder{
			&weightedFakeFinder{delay: 100 * time.Millisecond, weight: 10},
			&weightedFakeFinder{delay: 100 * time.Millisecond, weight: 10},
			&weightedFakeFinder{delay: 100 * time.Millisecond, weight: 10},
			&weightedFakeFinder{delay: 100 * time.Millisecond, weight: 10},
			&weightedFakeFinder{delay: 100 * time.Millisecond, weight: 10},
		}

		// Act
		start := time.Now()
		results := WeightedSemaphore("server-1", finders)
		elapsed := time.Since(start)

		// Assert
		require.Len(t, results, 5, "should return one result per finder")
		for _, result := range results {
			require.True(t, result.found)
		}
		require.GreaterOrEqual(t, elapsed, 200*time.Millisecond,
			"should take at least 200ms (weight-based batching prevents all running at once)")
		require.Less(t, elapsed, 300*time.Millisecond,
			"should allow 3 finders per batch (weight 10 * 3 = 30 capacity), not just 2")
	})
}

type weightedFakeFinder struct {
	delay  time.Duration
	weight int64
	calls  int
}

func (f *weightedFakeFinder) Find(resourceName string) bool {
	f.calls++
	time.Sleep(f.delay)
	return true
}

func (f *weightedFakeFinder) FindWithContext(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (f *weightedFakeFinder) FindWithError(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (f *weightedFakeFinder) Weight() int64 { return f.weight }
