package main

import (
	"context"
	"log/slog"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRedundant(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Arrange
		finders := []Finder{
			&contextAwareFinder{delay: 300 * time.Millisecond},
			&contextAwareFinder{delay: 100 * time.Millisecond},
			&contextAwareFinder{delay: 200 * time.Millisecond},
		}

		// Act
		start := time.Now()
		results := Redundant("server-1", finders)
		elapsed := time.Since(start)

		// Assert
		require.Len(t, results, 1, "should return only the first result, not all of them")
		require.True(t, results[0].found)
		require.Less(t, elapsed, 150*time.Millisecond,
			"should return as soon as the first finder responds")
	})
}

type contextAwareFinder struct {
	delay time.Duration
	calls int
}

func (f *contextAwareFinder) Find(_ string) bool { return true }
func (f *contextAwareFinder) Weight() int64       { return 1 }
func (f *contextAwareFinder) FindWithError(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (f *contextAwareFinder) FindWithContext(ctx context.Context, _ string) (bool, error) {
	f.calls++
	timer := time.NewTimer(f.delay)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

func (f *contextAwareFinder) LogValue() slog.Value {
	return slog.StringValue("fake")
}
