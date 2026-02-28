# A SOLID API

Start by having a look at the current code.

This project is an API that lists users from a database. Take a moment to explore the different files and understand how they relate to each other.

## Goal

Examine the codebase and identify what's wrong with the current design, using what you've learned about SOLID principles.

## Guiding questions

1. Open [`pkg/http/get_test.go`](./pkg/http/get_test.go). This is a unit test for the HTTP handler. What does it need to work? Does anything seem surprising for a test that's supposed to test HTTP behavior?

2. Trace the dependency chain starting from [`cmd/api/main.go`](./cmd/api/main.go). What object is passed from `main` to `Server`, and then from `Server` to the `user` package? What does this tell you about the coupling between layers?

3. Think about the **Single Responsibility Principle**: look at the [`pkg/user/`](./pkg/user/) package. How many responsibilities does it have? Should a "user" package know about SQL?

4. Think about the **Dependency Inversion Principle**: the HTTP handler (high-level) calls `user.List(s.db)` (low-level). Who depends on whom? Is that the right direction?

5. Imagine you need to switch from PostgreSQL to another storage system. How many files would you need to change? What does that tell you about the flexibility of this design?

## What you should take away

By the end of this exercise, you should be able to articulate:

- Why the current HTTP test needs SQL-level mocking (sqlmock) even though it's testing HTTP behavior
- That `*sql.DB` travels through every layer of the application, creating tight coupling
- Which SOLID principles are violated and why that matters

Take notes — you'll fix these issues in the next exercises!
