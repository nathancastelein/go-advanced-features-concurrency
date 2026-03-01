# Scatter Gather

The scatter gather is a concurrency pattern where computation is scattered, and then gathered.

To apply this pattern, we want to decorrelate the call to `Find`:

```
slog.Info("starting find", slog.Any("datacenter", finder))
found := finder.Find(resourceName)
```

And the response handling:

```
slog.Info("got result", slog.Any("datacenter", finder), slog.Bool("found", found))
```

To do so, we will use channels.

## First step

For the first step, we will modify a bit the response handling, to remove the datacenter information in the log:

Create a chan of bool, send response from `Find` into this chan.

Add a range loop to iterate on the chan, and handle the result by printing a log:

```
slog.Info("got result", slog.Bool("found", found))
```

## Second step

Let's now imagine we want to know the "origin" of each result, ie. the datacenter where the result comes from.

To do so, we will need to enhance a bit our result type.

Have a look at the [Result type](./result.go) provided in this package.

Then:
- Modify your chan to handle `Result` instead of `bool`.
- Create and send to the chan a proper `Result` with the result and the finder
- Add a new value in the `got result` log: `slog.Any("datacenter", result.datacenter)`

To test your code, run the test:

```bash
go test -run TestScatterGather -v
```

You can also run the application to see the logs:

```bash
go run . -action=scattergather
```