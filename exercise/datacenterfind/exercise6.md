# Semaphore

Semaphore pattern is a way to avoid sending too many jobs in the same time.

We relies on channels to do so and a lock/release lock mechanism.

Let's say we want to allow only two concurrent calls at the same time.

- Create a chan (chan type is not important) with the size 2.
- At the beginning of your goroutine, try to get a lock by sending data to your chan.
- At the end of your goroutine, read data from your chan to release the lock.


Test your code:

```bash
go run *go -action=semaphore
```

Expected output:

```
2024/06/29 23:56:36 INFO starting find datacenter=WAW
2024/06/29 23:56:36 INFO starting find datacenter=GRA
2024/06/29 23:56:36 INFO got result datacenter=WAW found=true
2024/06/29 23:56:36 INFO starting find datacenter=SBG
2024/06/29 23:56:36 INFO got result datacenter=GRA found=true
2024/06/29 23:56:36 INFO starting find datacenter=BHS
2024/06/29 23:56:36 INFO got result datacenter=SBG found=true
2024/06/29 23:56:36 INFO starting find datacenter=RBX
2024/06/29 23:56:36 INFO got result datacenter=RBX found=true
2024/06/29 23:56:36 INFO got result datacenter=BHS found=true
```

You program starts with two goroutines. Each time a result is found, it means one goroutine finished and released its lock, so a new one can start.