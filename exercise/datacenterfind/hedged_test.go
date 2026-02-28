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

type hedgedLogRecorder struct {
	mu      sync.Mutex
	entries []string
}

func (r *hedgedLogRecorder) Enabled(_ context.Context, _ slog.Level) bool { return true }
func (r *hedgedLogRecorder) WithAttrs(_ []slog.Attr) slog.Handler         { return r }
func (r *hedgedLogRecorder) WithGroup(_ string) slog.Handler               { return r }
func (r *hedgedLogRecorder) Handle(_ context.Context, record slog.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries = append(r.entries, record.Message)
	return nil
}

func (r *hedgedLogRecorder) count(msg string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	n := 0
	for _, e := range r.entries {
		if e == msg {
			n++
		}
	}
	return n
}

func TestHedged(t *testing.T) {
	t.Run("first finder fast enough", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			recorder := &hedgedLogRecorder{}
			slog.SetDefault(slog.New(recorder))

			// First finder responds in 50ms, well under the 75ms hedge timeout
			finders := []Finder{
				&contextAwareFinder{delay: 50 * time.Millisecond},
				&contextAwareFinder{delay: 200 * time.Millisecond},
				&contextAwareFinder{delay: 300 * time.Millisecond},
			}

			Hedged("server-1", finders)

			// Only the first finder should have been launched
			require.Equal(t, 1, recorder.count("launching find"),
				"should only launch the first finder when it responds fast enough")

			require.Equal(t, 1, recorder.count("got result"),
				"should have exactly one result")

			// Second and third finders should not have been called
			require.Equal(t, 0, finders[1].(*contextAwareFinder).calls,
				"second finder should not be called")
			require.Equal(t, 0, finders[2].(*contextAwareFinder).calls,
				"third finder should not be called")
		})
	})

	t.Run("fallback to second finder", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			recorder := &hedgedLogRecorder{}
			slog.SetDefault(slog.New(recorder))

			// First finder is slow (200ms), second responds in 50ms
			// After 75ms hedge timeout, second finder is launched and responds at 75+50=125ms
			finders := []Finder{
				&contextAwareFinder{delay: 200 * time.Millisecond},
				&contextAwareFinder{delay: 50 * time.Millisecond},
				&contextAwareFinder{delay: 300 * time.Millisecond},
			}

			start := time.Now()
			Hedged("server-1", finders)
			elapsed := time.Since(start)

			// Should complete at ~125ms (75ms wait + 50ms second finder), not 200ms
			require.Less(t, elapsed, 150*time.Millisecond,
				"should return from second finder before first finishes")

			// Two finders should have been launched, not three
			require.Equal(t, 2, recorder.count("launching find"),
				"should launch second finder after hedge timeout")

			require.Equal(t, 0, finders[2].(*contextAwareFinder).calls,
				"third finder should not be called")
		})
	})
}
