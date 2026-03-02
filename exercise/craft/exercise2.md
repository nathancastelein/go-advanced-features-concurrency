# Decoupling with Interfaces

In the previous exercise, you identified that all layers are tightly coupled through `*sql.DB`. The HTTP handler test needs sqlmock, the `user` package mixes business logic with SQL, and changing the storage would require modifying almost every file.

Let's fix this using interfaces.

## Goal

Refactor the codebase so that the HTTP handler no longer depends on `*sql.DB` or on the SQL logic. Use an interface to inverse the dependency.

## What you need to achieve

1. **The `user` package should only contain business definitions**: the `User` struct and an interface describing how to retrieve users. No SQL, no `database/sql` import.

2. **SQL logic should live in its own package**: create a dedicated package (for example `storage`) that contains the struct and methods responsible for querying the database. This package should implement the interface defined in `user`.

3. **The HTTP `Server` should depend on the interface, not on `*sql.DB`**: the handler doesn't need to know *how* users are fetched, only *that* they can be fetched.

4. **`main.go` should wire everything together**: this is the only place that knows about both the HTTP server and the concrete storage implementation.

## Validation

When you're done:

- `go build ./...` should compile without errors
- `go test ./pkg/storage/...` should pass (the SQL test with sqlmock)
- The `pkg/http/` package should **not** import `database/sql`

## Step by step

1. In `pkg/user/`, define a `Lister` interface with a single method: `List() ([]User, error)`. Keep it alongside the `User` struct. Remove all SQL-related code from this package.

2. Create a `pkg/storage/` package. Move the SQL logic there into a struct (e.g. `UserSQL`) with a `List() ([]user.User, error)` method. In Go, interface implementation is implicit: if the method signature matches, the interface is satisfied automatically.

3. You can make the constructor return the interface type to verify at compile time that the implementation is correct:

```go
func NewUserSQL(db *sql.DB) user.Lister {
    return &UserSQL{db: db}
}
```

4. Update `pkg/http/` so that `Server` takes a `user.Lister` instead of `*sql.DB`.

5. Update `cmd/api/main.go` to create the storage and pass it to the HTTP server.
