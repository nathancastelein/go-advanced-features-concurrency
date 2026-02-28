# Exercise 1

Open the [TimerTicker function](./timerticker/timerticker.go).

TimerTicker is a simple function that you will need to write.

Using [time.Ticker](https://pkg.go.dev/time#Ticker) and [time.Timer](https://pkg.go.dev/time#Timer), write a function that:

- Starts a timer with the given input timerDuration
- Starts a ticker with the given input tickerDuration
- Using an infinite loop and a select statement, listen on both ticker & timer channels
- At each tick, use `fmt.Println` to write "Hello from ticker"
- At the end of the timer, use `fmt.Println` to write "Time to say goodbye"

Use `go test .` to test your function.

You can have a look on [TimerTicker test function](./timerticker/timerticker_test.go) to understand the expected behavior.
