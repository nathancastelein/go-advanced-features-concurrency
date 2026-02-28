package main

import (
	"fmt"
	"sync"
)

var x = 0
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//mu.Lock()
			//defer mu.Unlock()
			defer wg.Done()
			x = x + 1
		}()
	}
	wg.Wait()
	fmt.Printf("X = %d\n", x)
}
