# Decoupling with Interfaces

In the previous exercise, you identified that all layers of this application are tightly coupled through `*sql.DB`. The HTTP handler test needs sqlmock, the `user` package mixes business logic with SQL, and changing the storage would require modifying almost every file.

Let's fix this using interfaces and the SOLID principles we've seen in the slides.

## Goal

Refactor the codebase so that the HTTP handler no longer depends on `*sql.DB` or on the SQL logic in the `user` package. Use an interface to inverse the dependency.

## What you need to achieve

Here are your objectives — it's up to you to figure out how to get there:

1. **The `user` package should only contain business definitions**: the `User` struct and an interface describing how to retrieve users. No SQL, no `database/sql` import.

2. **SQL logic should live in its own package**: create a dedicated package (for example `storage`) that contains the struct and methods responsible for querying the database. This package should implement the interface defined in `user`.

3. **The HTTP `Server` should depend on the interface, not on `*sql.DB`**: the handler doesn't need to know *how* users are fetched, only *that* they can be fetched.

4. **`main.go` should wire everything together**: this is the only place that knows about both the HTTP server and the concrete storage implementation.

## Which SOLID principles are you applying?

As you refactor, think about which principles guide each change:

- **Single Responsibility Principle (SRP)**: the `user` package was doing two things — defining business objects and querying the database. After refactoring, each package has one job.
- **Dependency Inversion Principle (DIP)**: the HTTP handler (high-level module) will depend on an abstraction (the interface), not on the SQL implementation (low-level module).
- **Interface Segregation Principle (ISP)**: the interface should be small and focused — only the method(s) the HTTP handler actually needs.

## Validation

When you're done:

- `go build ./...` should compile without errors
- `go test ./pkg/storage/...` should pass (the SQL test with sqlmock)
- The `pkg/http/` package should **not** import `database/sql`

## Hints

<details>
<summary>Hint 1: the interface</summary>

Define a `Lister` interface with a single method: `List() ([]User, error)`. Put it in the `user` package alongside the `User` struct.

</details>

<details>
<summary>Hint 2: implicit implementation</summary>

In Go, interface implementation is implicit. If your storage struct has a `List() ([]User, error)` method, it automatically satisfies the `Lister` interface. No need for `implements` keywords.

</details>

<details>
<summary>Hint 3: verify implementation at compile time</summary>

You can make the constructor return the interface type instead of the concrete type to verify that the implementation is correct at compile time:

```go
func NewUserSQL(db *sql.DB) user.Lister {
    return &UserSQL{db: db}
}
```

</details>
