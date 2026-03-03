package timerticker

import "time"

func TimerTicker(timerDuration time.Duration, tickerDuration time.Duration) int {
	timer := time.NewTimer(timerDuration)
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()
	defer timer.Stop()

	ticks := 0
	for {
		select {
		case <-timer.C:
			return ticks
		case <-ticker.C:
			ticks++
		}
	}
}
