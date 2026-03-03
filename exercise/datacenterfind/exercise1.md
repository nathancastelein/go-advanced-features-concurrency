# WaitGroup

Open [waitgroup.go](./waitgroup.go) file.

Using `sync.WaitGroup` ([https://pkg.go.dev/sync#WaitGroup](https://pkg.go.dev/sync#WaitGroup)), add concurrency so the calls to `finder.Find` are performed simultaneously.

To test your development, run the test:

```bash
go test -run TestWaitGroup -v
```

You can also run the application to see the logs:

```bash
go run . -action=waitgroup
```