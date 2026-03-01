# Semaphore

Semaphore pattern is a way to avoid sending too many jobs in the same time.

We relies on channels to do so and a lock/release lock mechanism.

Let's say we want to allow only two concurrent calls at the same time.

- Create a chan (chan type is not important) with the size 2.
- At the beginning of your goroutine, try to get a lock by sending data to your chan.
- At the end of your goroutine, read data from your chan to release the lock.


To test your code, run the test:

```bash
go test -run TestSemaphore -v
```

You can also run the application to see the logs:

```bash
go run . -action=semaphore
```