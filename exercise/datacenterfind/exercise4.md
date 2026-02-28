# ErrGroup

Let's now using a function with an error. Experimental sync package provides the `ErrGroup` ([https://pkg.go.dev/golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)) structure for this kind of purpose.

ErrGroup is close to WaitGroup in its behaviour. It's an helper to manage concurrency with functions which can return an error.

Let's use it by fixing the panic in [errgroup.go](./errgroup.go).

The goal of this exercise is to change the `sync.WaitGroup` to a `ErrGroup`.

- Creates a new `ErrGroup` by using [errgroup.WithContext](https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext).
- Replace the `go func()` with the proper ErrGroup `Go` method.
- Replace the `wg.Wait` to use the ErrGroup `Wait` method.
- Handle the error from ErrGroup `Wait` method with an error log: `		slog.Error("an error occured", slog.String("error", err.Error()))`
- Remove the `sync.WaitGroup`.

Test your code:

```bash
go run *go -action=errgroup
```

Expected output:

```
2024/06/29 23:29:06 INFO deadline exceeded finder=GRA
2024/06/29 23:29:06 INFO deadline exceeded finder=SBG
2024/06/29 23:29:06 INFO deadline exceeded finder=BHS
2024/06/29 23:29:06 ERROR an error occured error="something went wrong"
```
