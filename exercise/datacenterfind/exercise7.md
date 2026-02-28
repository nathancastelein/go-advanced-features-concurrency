# Weighted Semaphore

Let's now rewriter our homemade semaphore with a [Weighted Semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore#NewWeighted).

- Create a new weighted semaphore with a weight of 30
- Use `Acquire` method to acquire a lock. You can use `finder.Weight()` to get the weight of your actual finder.
- Use `Release` method to release the lock.

Test your code:

```bash
go run *go -action=wsemaphore
```