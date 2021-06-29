package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Threads", runtime.GOMAXPROCS(-1)) // All cores e.g. 4
	runtime.GOMAXPROCS(1)
	fmt.Println("Threads",runtime.GOMAXPROCS(1)) // Now using 1 core

	fmt.Println("Starting")
	go sayHello()
	time.Sleep(1 * time.Second) // Sleeps for 1 second - Adding sleep is not a good practice
	fmt.Println("Ending")
}

func sayHello() {
	fmt.Println("Mid")
}
