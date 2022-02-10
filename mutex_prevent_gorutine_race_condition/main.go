package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	counter := 0
	var gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//Lock from Read and Write counter variable
			mu.Lock()

			v := counter //Read the value
			runtime.Gosched()
			v++          //Write the value
			counter = v

			mu.Unlock() //Unlock from Read and Write counter variable

			wg.Done()
		}()
		fmt.Println("Goroutine: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
