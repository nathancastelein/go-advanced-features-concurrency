# A SOLID API

This project is an API that lists users from a database. Take a moment to explore the different files and understand how they relate to each other.

## Goal

Examine the codebase and identify design issues using the SOLID principles.

## SOLID principles recap

- **S — Single Responsibility**: a module should have one, and only one, reason to change.
- **O — Open/Closed**: software entities should be open for extension, but closed for modification.
- **L — Liskov Substitution**: objects should be replaceable with instances of their subtypes without altering correctness.
- **I — Interface Segregation**: clients should not be forced to depend on methods they do not use.
- **D — Dependency Inversion**: high-level modules should not depend on low-level modules. Both should depend on abstractions.

## Guiding questions

1. Open `pkg/http/get_test.go`. What does this test need to run? Is that surprising for a test that's supposed to test HTTP behavior?

2. Starting from `cmd/api/main.go`, trace what is passed from layer to layer. What does this tell you about the coupling?

3. Look at `pkg/user/`. How many responsibilities does this package have?

4. If you needed to switch from PostgreSQL to another storage, how many files would you need to change?

5. Which SOLID principles are violated here?

Take notes — you'll fix these issues in the next exercises!
