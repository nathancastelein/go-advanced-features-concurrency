package main

import (
	"context"
	"log/slog"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

type logEntry struct {
	message string
}

type logRecorder struct {
	mu      sync.Mutex
	entries []logEntry
}

func (r *logRecorder) Enabled(_ context.Context, _ slog.Level) bool { return true }
func (r *logRecorder) WithAttrs(_ []slog.Attr) slog.Handler           { return r }
func (r *logRecorder) WithGroup(_ string) slog.Handler                 { return r }
func (r *logRecorder) Handle(_ context.Context, record slog.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries = append(r.entries, logEntry{message: record.Message})
	return nil
}

func TestScatterGather(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		recorder := &logRecorder{}
		slog.SetDefault(slog.New(recorder))

		finders := []Finder{
			&fakeFinder{delay: 100 * time.Millisecond},
			&fakeFinder{delay: 200 * time.Millisecond},
			&fakeFinder{delay: 300 * time.Millisecond},
		}

		start := time.Now()
		ScatterGather("server-1", finders)
		elapsed := time.Since(start)

		// All finders must have been called
		for i, f := range finders {
			require.Equal(t, 1, f.(*fakeFinder).calls, "finder %d should have been called once", i)
		}

		// With concurrency, total time should be ~300ms (slowest), not 600ms (sum)
		require.Less(t, elapsed, 400*time.Millisecond, "finders should run concurrently, not sequentially")

		// Verify we have the expected log pattern: 3 "starting find" + 3 "got result"
		var startCount, resultCount int
		for _, e := range recorder.entries {
			switch e.message {
			case "starting find":
				startCount++
			case "got result":
				resultCount++
			}
		}
		require.Equal(t, 3, startCount, "should have 3 'starting find' logs")
		require.Equal(t, 3, resultCount, "should have 3 'got result' logs")
	})
}
