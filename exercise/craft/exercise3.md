# Testing with Doubles

After the refactoring in Exercise 2, the HTTP handler now depends on an interface (`user.Lister`) instead of directly on `*sql.DB`. This is a big improvement — but look at the HTTP handler test. It hasn't been updated yet.

## Goal

Leverage the interface to write simpler, faster unit tests for the HTTP handler. Replace the sqlmock dependency with a test double.

## The problem

Open [`pkg/http/get_test.go`](./pkg/http/get_test.go).

This test still sets up a SQL mock (`sqlmock.New()`), configures expected queries, and passes a `*sql.DB` to the `Server`. But now that the `Server` depends on a `user.Lister` interface, do we still need all of this SQL machinery to test an HTTP handler?

## What you need to achieve

1. **Create a stub** that implements the `user.Lister` interface. This stub should return fixed, predetermined data — no database, no SQL, just hardcoded return values.

2. **Rewrite the HTTP handler test** to use your stub instead of sqlmock. The test should be much simpler: create a stub, create a server with the stub, call the handler, assert the response.

3. **Clean up**: after your changes, the HTTP test file should not import `go-sqlmock` or `database/sql` anymore.

4. **Don't forget `main.go`**: if you haven't already, update [`cmd/api/main.go`](./cmd/api/main.go) to wire your concrete storage implementation into the HTTP server through the interface.

## Reflection

Once your tests pass, think about these questions:

- **Stub vs mock**: a stub returns fixed data. A mock verifies that specific methods were called with specific arguments. Which one did you use? Which one does sqlmock provide?
- **Where should sqlmock stay?** Your storage package tests still use sqlmock — and that's correct! Those tests actually verify SQL queries. The HTTP handler tests shouldn't care about SQL.
- **Speed and reliability**: how much simpler is the new test? What would happen if the SQL schema changed — which tests would break?

## Validation

When you're done:

- `go test ./pkg/http/...` passes
- `go test ./pkg/storage/...` passes
- The HTTP test file does **not** import `go-sqlmock` or `database/sql`
- `go build ./...` compiles (including the updated `main.go`)

## Hints

<details>
<summary>Hint 1: what's a stub?</summary>

A stub is a simple struct with methods that return hardcoded data. No logic, no side effects — just return values.

```go
type MyStub struct{}

func (s *MyStub) List() ([]user.User, error) {
    return []user.User{{Firstname: "Test", Lastname: "User"}}, nil
}
```

</details>

<details>
<summary>Hint 2: compile-time interface check</summary>

You can verify at compile time that your stub correctly implements the interface:

```go
var _ user.Lister = &MyStub{}
```

If the stub doesn't implement the interface, the compiler will tell you.

</details>

Well done — you've applied SOLID principles to decouple layers, and now you can see the concrete benefit: simpler, faster, more focused tests!
