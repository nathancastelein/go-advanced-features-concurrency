package timerticker

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTickTimer(t *testing.T) {
	// Arrange
	orig := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	defer func() {
		os.Stdout = orig
	}()
	expectedOutput := strings.Repeat("Hello from ticker\n", 5) + "Time to say goodbye\n"

	// Act
	TimerTicker(55*time.Millisecond, 10*time.Millisecond)

	// Assert
	w.Close()
	os.Stdout = orig
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	require.NotNil(t, out)
	require.Equal(t, expectedOutput, string(out))
}
