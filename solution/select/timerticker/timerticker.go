package timerticker

import (
	"fmt"
	"time"
)

func TimerTicker(timerDuration time.Duration, tickerDuration time.Duration) {
	timer := time.NewTimer(timerDuration)
	ticker := time.NewTicker(tickerDuration)

	for {
		select {
		case <-timer.C:
			fmt.Println("Time to say goodbye")
			ticker.Stop()
			timer.Stop()
			return
		case <-ticker.C:
			fmt.Println("Hello from ticker")
		}
	}
}
