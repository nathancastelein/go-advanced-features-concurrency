package main

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestErrGroup(t *testing.T) {
	t.Run("all finders succeed", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			// Arrange
			finders := []Finder{
				&errorFinder{delay: 100 * time.Millisecond},
				&errorFinder{delay: 200 * time.Millisecond},
				&errorFinder{delay: 300 * time.Millisecond},
			}

			// Act
			results, err := ErrGroup("server-1", finders)

			// Assert
			require.NoError(t, err)
			require.Len(t, results, 3, "should return one result per finder")
			for _, result := range results {
				require.True(t, result.found)
			}
		})
	})

	t.Run("one finder returns an error", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			// Arrange
			finders := []Finder{
				&errorFinder{delay: 100 * time.Millisecond},
				&errorFinder{delay: 200 * time.Millisecond, err: errors.New("something went wrong")},
				&errorFinder{delay: 300 * time.Millisecond},
			}

			// Act
			results, err := ErrGroup("server-1", finders)

			// Assert
			require.Error(t, err, "should return the error from the failing finder")
			require.Nil(t, results, "should not return results when an error occurs")
		})
	})
}

type errorFinder struct {
	delay time.Duration
	err   error
}

func (f *errorFinder) Find(_ string) bool { return true }
func (f *errorFinder) FindWithContext(_ context.Context, _ string) (bool, error) {
	return false, nil
}
func (f *errorFinder) Weight() int64 { return 1 }

func (f *errorFinder) FindWithError(ctx context.Context, _ string) (bool, error) {
	if f.err != nil {
		return false, f.err
	}
	timer := time.NewTimer(f.delay)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}
