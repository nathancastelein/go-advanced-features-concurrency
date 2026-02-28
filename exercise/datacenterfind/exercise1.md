# WaitGroup

Open [waitgroup.go](./waitgroup.go) file.

Using `sync.WaitGroup` ([https://pkg.go.dev/sync#WaitGroup](https://pkg.go.dev/sync#WaitGroup)), add concurrency so the calls to `finder.Find` are performed simultaneously.

To test your development, you can run the program:

```bash
go run *go -action=waitgroup
```

Expected output: three `starting find` logs, followed by three `got result`.

```
2024/06/29 22:16:58 INFO starting find datacenter=BHS
2024/06/29 22:16:58 INFO starting find datacenter=SBG
2024/06/29 22:16:58 INFO starting find datacenter=GRA
2024/06/29 22:16:58 INFO got result datacenter=SBG found=true
2024/06/29 22:16:58 INFO got result datacenter=GRA found=true
2024/06/29 22:16:58 INFO got result datacenter=BHS found=true
```