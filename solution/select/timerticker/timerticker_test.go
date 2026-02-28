package timerticker

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTickTimer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ticks := TimerTicker(55*time.Millisecond, 10*time.Millisecond)
		require.Equal(t, 5, ticks)
	})
}
