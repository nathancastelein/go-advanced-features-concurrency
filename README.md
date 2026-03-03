# Go Advanced Features & Concurrency - Exercises

This repository contains the exercises for Module 3 of the Go Training course.

## Prerequisites

- Go 1.25 or later

## Repository Structure

```
go-advanced-features-concurrency/
├── exercise/           # Exercise stubs - implement these!
│   ├── generics/       # Generic functions and structs
│   ├── interfaces/     # Interface patterns
│   ├── craft/          # Software craft / clean architecture
│   ├── concurrency/    # Basic concurrency (waitgroup, channel)
│   ├── sync/           # Sync primitives (mutex, once)
│   ├── select/         # Select statement exercises
│   └── datacenterfind/ # Progressive concurrency patterns
└── solution/           # Complete solutions
    ├── generics/
    ├── interfaces/
    ├── craft/
    ├── concurrency/
    ├── sync/
    ├── select/
    └── datacenterfind/
```

## Exercise Order

### Part 1: Advanced Features

| # | Exercise | File to open | Description |
|---|----------|-------------|-------------|
| 1 | Generics: Contains | `exercise/generics/exercise1.md` | Rewrite Contains using generics |
| 2 | Generics: Queue | `exercise/generics/exercise2.md` | Rewrite Queue using generics |
| 3 | Interfaces: Pricer | `exercise/interfaces/store/exercice1.md` | Introduce a Pricer interface |
| 4 | Interfaces: Test Double | `exercise/interfaces/store/exercice2.md` | Write a StubPricer for unit tests |

### Part 2: Software Craft

| # | Exercise | File to open | Description |
|---|----------|-------------|-------------|
| 5 | Craft: Exercise 1 | `exercise/craft/exercise1.md` | Refactoring step 1 - Extract SQL |
| 6 | Craft: Exercise 2 | `exercise/craft/exercise2.md` | Refactoring step 2 - Introduce interface |
| 7 | Craft: Exercise 3 | `exercise/craft/exercise3.md` | Refactoring step 3 - Dependency injection |

### Part 3: Basic Concurrency

| # | Exercise | File to open | Description |
|---|----------|-------------|-------------|
| 8 | Concurrency | `exercise/concurrency/exercise1.md` | Add concurrency with WaitGroup and channels |

### Part 4: Sync Primitives

| # | Exercise | File to open | Description |
|---|----------|-------------|-------------|
| 9 | Mutex | `exercise/sync/mutex/exercise1.md` | Add mutex to prevent race condition |
| 10 | sync.Once | `exercise/sync/once/exercise1.md` | Implement singleton with sync.Once |

### Part 5: Select

| # | Exercise | File to open | Description |
|---|----------|-------------|-------------|
| 11 | Timer & Ticker | `exercise/select/exercise1.md` | Use select with Timer and Ticker |

### Part 6: Concurrency Patterns (Datacenter Finder)

| # | Exercise | File to open | Description |
|---|----------|-------------|-------------|
| 12 | WaitGroup | `exercise/datacenterfind/exercise1.md` | Add basic concurrency |
| 13 | Scatter-Gather | `exercise/datacenterfind/exercise2.md` | Implement scatter-gather with channels |
| 14 | Redundant Requests | `exercise/datacenterfind/exercise3.md` | Use context cancellation |
| 15 | ErrGroup | `exercise/datacenterfind/exercise4.md` | Handle errors with errgroup |
| 16 | Hedged Requests | `exercise/datacenterfind/exercise5.md` | Implement hedged request pattern |
| 17 | Semaphore | `exercise/datacenterfind/exercise6.md` | Limit concurrency with semaphore |
| 18 | Weighted Semaphore | `exercise/datacenterfind/exercise7.md` | Use weighted semaphore from x/sync |

## Running Tests

### Test your exercise implementation

```bash
# Test a specific exercise
go test ./exercise/generics/...

# Test all exercises (will fail until implemented)
go test ./exercise/...
```

### Verify against solutions

```bash
# All solution tests should pass
go test ./solution/...

# Run with race detector
go test -race ./solution/...
```

## Go Version

This module requires Go 1.25 or later. Check your version:

```bash
go version
```
