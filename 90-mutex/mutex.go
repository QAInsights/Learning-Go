package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var c int = 0
var m = sync.RWMutex{}

func main() {

	// To test race conditions, send `go run .\90-mutex\mutex.go -race`
	fmt.Println("Mutex demo")

	for j := 0; j<10; j++{

		wg.Add(2)
		m.Lock()
		go incrementer()

		m.RLock()
		go sayHello()
	}

	wg.Wait()
}

func sayHello(){
	fmt.Println("Hello World", c)
	m.RUnlock()
	wg.Done()
}

func incrementer(){
	c++
	fmt.Println("Incremented to ",c)
	m.Unlock()
	wg.Done()
}