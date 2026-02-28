# Hedged requests

Hedged requests is a concurrency pattern where we want to send a first request, and if the answer is not coming in a given time, then we send the second request, etc.

To do so, we will use a `time.Timer` and play with a `select` statement.

- Change the size of your chan to 0, as we will need only one result
- Before the goroutine creation, create a `time.Timer` with a duration of `75 milliseconds`.
- After your goroutine, add a `select` statement to select between your result chan and the `timer.C` chan.
- Don't forget to exit your function at the first result, and to cancel context and close your results chan

Test your code:

```bash
go run *go -action=hedged
```

Expected output:

```
2024/07/08 22:41:35 INFO launching find datacenter=SBG
2024/07/08 22:41:35 INFO launching find datacenter=GRA
2024/07/08 22:41:35 INFO got result found=true
```

The first FindWithContext took too much time, so another FindWithContext has been launched. As we finally had a result, there are no other FindWithContext launched.