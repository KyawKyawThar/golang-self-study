package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Loop GR:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
