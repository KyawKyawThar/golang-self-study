package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Using Atomic to prevent race condition
func main() {

	fmt.Println("first Number of Goroutine: ", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Loop of Goroutine: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Last of Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
