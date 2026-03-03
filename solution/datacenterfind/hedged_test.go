package main

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHedged(t *testing.T) {
	t.Run("first finder fast enough", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			// Arrange: first finder responds in 50ms, well under the 75ms hedge timeout
			finders := []Finder{
				&contextAwareFinder{delay: 50 * time.Millisecond},
				&contextAwareFinder{delay: 200 * time.Millisecond},
				&contextAwareFinder{delay: 300 * time.Millisecond},
			}

			// Act
			results := Hedged("server-1", finders)

			// Assert
			require.Len(t, results, 1, "should return only the first result")
			require.True(t, results[0].found)
			require.Equal(t, 0, finders[1].(*contextAwareFinder).calls,
				"second finder should not be called")
			require.Equal(t, 0, finders[2].(*contextAwareFinder).calls,
				"third finder should not be called")
		})
	})

	t.Run("fallback to second finder", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			// Arrange: first finder is slow (200ms), second responds in 50ms
			// After 75ms hedge timeout, second finder is launched and responds at 75+50=125ms
			finders := []Finder{
				&contextAwareFinder{delay: 200 * time.Millisecond},
				&contextAwareFinder{delay: 50 * time.Millisecond},
				&contextAwareFinder{delay: 300 * time.Millisecond},
			}

			// Act
			start := time.Now()
			results := Hedged("server-1", finders)
			elapsed := time.Since(start)

			// Assert
			require.Len(t, results, 1, "should return only the first result")
			require.True(t, results[0].found)
			require.Less(t, elapsed, 150*time.Millisecond,
				"should return from second finder before first finishes")
			require.Equal(t, 0, finders[2].(*contextAwareFinder).calls,
				"third finder should not be called")
		})
	})
}
