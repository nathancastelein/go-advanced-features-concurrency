package main

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestScatterGather(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Arrange
		finders := []Finder{
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 200 * time.Millisecond},
			&fakeFinder{delay: 300 * time.Millisecond},
		}

		// Act
		start := time.Now()
		results := ScatterGather("server-1", finders)
		elapsed := time.Since(start)

		// Assert
		require.Len(t, results, 3, "should return one result per finder")
		for _, result := range results {
			require.True(t, result.found)
		}
		require.Less(t, elapsed, 400*time.Millisecond, "finders should run concurrently, not sequentially")
	})
}
