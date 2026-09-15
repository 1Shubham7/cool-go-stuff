package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int = 0
	var mu sync.Mutex

	var wg sync.WaitGroup
	
	for i := range 100 {	
		wg.Add(1)
		go func() { 
			defer wg.Done()
			mu.Lock()
			sum = sum + i
			mu.Unlock()
		}()
	}
	wg.Wait()


	fmt.Print(sum)
}