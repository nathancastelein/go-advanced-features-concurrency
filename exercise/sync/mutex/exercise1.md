# Mutex

Open [main.go](main.go) file.

This file proposes a `Storage` class to store information into a map.

Try to run the program:

```
go run main.go
```

Does everything look fine? Add now the `-race` flag.

```
go run -race main.go
```

There's a race condition!

Use a [sync.Mutex](https://pkg.go.dev/sync#Mutex) to fix this.

Test your code by running the code with a `-race`:

```
go run -race main.go
```