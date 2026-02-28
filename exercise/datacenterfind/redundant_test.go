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

type redundantLogRecorder struct {
	mu      sync.Mutex
	entries []string
}

func (r *redundantLogRecorder) Enabled(_ context.Context, _ slog.Level) bool { return true }
func (r *redundantLogRecorder) WithAttrs(_ []slog.Attr) slog.Handler         { return r }
func (r *redundantLogRecorder) WithGroup(_ string) slog.Handler               { return r }
func (r *redundantLogRecorder) Handle(_ context.Context, record slog.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries = append(r.entries, record.Message)
	return nil
}

func TestRedundant(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		recorder := &redundantLogRecorder{}
		slog.SetDefault(slog.New(recorder))

		finders := []Finder{
			&contextAwareFinder{delay: 300 * time.Millisecond},
			&contextAwareFinder{delay: 100 * time.Millisecond},
			&contextAwareFinder{delay: 200 * time.Millisecond},
		}

		start := time.Now()
		Redundant("server-1", finders)
		elapsed := time.Since(start)

		// Should complete in ~100ms (fastest finder), not 300ms (slowest)
		require.Less(t, elapsed, 150*time.Millisecond,
			"should return as soon as the first finder responds")

		// Should have only 1 "got result" log (first answer only)
		var resultCount int
		for _, msg := range recorder.entries {
			if msg == "got result" {
				resultCount++
			}
		}
		require.Equal(t, 1, resultCount,
			"should only process the first result, not all of them")
	})
}
