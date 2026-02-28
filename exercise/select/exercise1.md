# Exercise 1

Open the [TimerTicker function](./timerticker/timerticker.go).

TimerTicker is a simple function that you will need to write.

Using [time.Ticker](https://pkg.go.dev/time#Ticker) and [time.Timer](https://pkg.go.dev/time#Timer), write a function that:

- Starts a timer with the given input timerDuration
- Starts a ticker with the given input tickerDuration
- Using an infinite loop and a select statement, listen on both ticker & timer channels
- At each tick, increment a counter
- When the timer fires, return the number of ticks that occurred

Use `go test .` to test your function.