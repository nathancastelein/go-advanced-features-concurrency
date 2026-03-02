# Testing with Doubles

After the refactoring in Exercise 2, the HTTP handler now depends on an interface (`user.Lister`) instead of `*sql.DB`. But the HTTP handler test hasn't been updated yet — it still uses sqlmock.

## Goal

Replace the sqlmock dependency in the HTTP handler test with a simple test double (stub).

## Validation

When you're done:

- `go test ./pkg/http/...` passes
- `go test ./pkg/storage/...` passes
- The HTTP test file does **not** import `go-sqlmock` or `database/sql`
- `go build ./...` compiles

## Step by step

1. Create a stub struct that implements `user.Lister`. A stub simply returns hardcoded data:

```go
type stubLister struct{}

func (s *stubLister) List() ([]user.User, error) {
    return []user.User{{Firstname: "Test", Lastname: "User"}}, nil
}
```

2. You can verify at compile time that your stub implements the interface:

```go
var _ user.Lister = &stubLister{}
```

3. Rewrite the HTTP handler test to use your stub instead of sqlmock. The test becomes: create a stub, create a server with the stub, call the handler, assert the response.

4. Clean up: remove the `go-sqlmock` and `database/sql` imports from the HTTP test file. sqlmock should only remain in `pkg/storage/` tests where it actually verifies SQL queries.

5. If you haven't already, update `cmd/api/main.go` to wire the concrete storage into the HTTP server through the interface.
