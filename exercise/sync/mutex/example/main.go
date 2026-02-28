package main

import (
	"fmt"
	"sync"
)

var x = 0
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Go(func() {
			//mu.Lock()
			//defer mu.Unlock()
			x += 1
		})
	}
	wg.Wait()
	fmt.Printf("X = %d\n", x)
}
