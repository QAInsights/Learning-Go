package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0

func main() {

	// Spinning up 10 go routines without sleep
	// there will NOT be any sync across routines

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go greetHello()
	}

	wg.Wait()
}

func greetHello() {
	fmt.Println("Hello World!")
	wg.Done()
}
