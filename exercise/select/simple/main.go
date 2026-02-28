package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch1 <- "goroutine 1"
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch2 <- "goroutine 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Printf("Hello from %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("Hello from %s\n", msg)
	}
}
