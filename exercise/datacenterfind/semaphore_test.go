package main

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSemaphore(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Arrange: 5 finders each taking 100ms, max 2 parallel
		// With semaphore (max 2): ceil(5/2) * 100ms = 300ms
		// Without semaphore (fully concurrent): 100ms
		// Sequential: 500ms
		finders := []Finder{
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 100 * time.Millisecond},
		}

		// Act
		start := time.Now()
		results := Semaphore("server-1", finders)
		elapsed := time.Since(start)

		// Assert
		require.Len(t, results, 5, "should return one result per finder")
		for _, result := range results {
			require.True(t, result.found)
		}
		require.GreaterOrEqual(t, elapsed, 300*time.Millisecond,
			"should take at least 300ms with max 2 parallel (semaphore limits concurrency)")
		require.Less(t, elapsed, 500*time.Millisecond,
			"should not run sequentially (500ms), semaphore allows some parallelism")
	})
}
