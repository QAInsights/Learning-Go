package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// -1 will set the value and doesn't change or go back
	// Now using 1 thread, setting the threads to higher value will cause overhead

	fmt.Println("Threads", runtime.GOMAXPROCS(-1)) // All threads e.g. 4
	runtime.GOMAXPROCS(1)
	fmt.Println("Threads",runtime.GOMAXPROCS(-1))

	fmt.Println("Starting")
	go sayHello()
	time.Sleep(1 * time.Second) // Sleeps for 1 second - Adding sleep is not a good practice
	fmt.Println("Ending")
}

func sayHello() {
	fmt.Println("Mid")
}
