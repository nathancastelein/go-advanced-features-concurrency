# ErrGroup

Let's now using a function with an error. Experimental sync package provides the `ErrGroup` ([https://pkg.go.dev/golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)) structure for this kind of purpose.

ErrGroup is close to WaitGroup in its behaviour. It's an helper to manage concurrency with functions which can return an error.

Let's use it by fixing the panic in [errgroup.go](./errgroup.go).

The goal of this exercise is to change the `sync.WaitGroup` to a `ErrGroup`.

- Creates a new `ErrGroup` by using [errgroup.WithContext](https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext).
- Replace the `go func()` with the proper ErrGroup `Go` method.
- Replace the `wg.Wait` to use the ErrGroup `Wait` method.
- Handle the error from ErrGroup `Wait` method by returning it.
- Remove the `sync.WaitGroup`.

To test your code, run the test:

```bash
go test -run TestErrGroup -v
```

You can also run the application to see the logs:

```bash
go run . -action=errgroup
```
