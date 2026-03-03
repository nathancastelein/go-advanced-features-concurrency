# Redundant

Let's now imagine another use case: you just need the first result of all your calls, wherever it comes from.

To do so, we will use the redundant pattern: starting all the calls concurrently, wait for the first result and cancel all other running jobs.

The [context.Context structure](https://pkg.go.dev/context) is designed for this purpose, thanks to the [WithCancel](https://pkg.go.dev/context#WithCancel) mathod.

To understand how cancelation is handled, you can have a look on the `FindWithContext` method definition in [finder.go](finder.go).

Open the [redundant.go](./redundant.go) file:

- Create a new cancelable context from `context.Background()`
- After the `starting find` log, call the function `finder.FindWithContext` with your properly created context
- Send the result in your chan, only if the function returned no errors
- Stop the function after the first result, then call the cancel method
- Don't forget to close the channel

To test your code, run the test:

```bash
go test -run TestRedundant -v
```

You can also run the application to see the logs:

```bash
go run . -action=redundant
```