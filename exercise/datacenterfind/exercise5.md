# Hedged requests

Hedged requests is a concurrency pattern where we want to send a first request, and if the answer is not coming in a given time, then we send the second request, etc.

To do so, we will use a `time.Timer` and play with a `select` statement.

- Change the size of your chan to 0, as we will need only one result
- Before the goroutine creation, create a `time.Timer` with a duration of `75 milliseconds`.
- After your goroutine, add a `select` statement to select between your result chan and the `timer.C` chan.
- Don't forget to exit your function at the first result, and to cancel context and close your results chan

To test your code, run the test:

```bash
go test -run TestHedged -v
```

You can also run the application to see the logs:

```bash
go run . -action=hedged
```