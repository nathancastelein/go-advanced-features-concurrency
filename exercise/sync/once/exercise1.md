# sync.Once

Open [database.go](./database.go).

Fix the code with [sync.Once](https://pkg.go.dev/sync#Once) to ensure the database will be initialized only one time.

To test your code, you can run it:

```
go run *go
```

Expected output: only one log `connecting to database` then three logs with the same connection id.

```
2024/06/30 15:19:04 INFO connecting to database
2024/06/30 15:19:04 INFO got database connection! database=&{id:06e802ad-02d8-407d-b770-661e01b5ce5c}
2024/06/30 15:19:04 INFO got database connection! database=&{id:06e802ad-02d8-407d-b770-661e01b5ce5c}
2024/06/30 15:19:04 INFO got database connection! database=&{id:06e802ad-02d8-407d-b770-661e01b5ce5c}
```