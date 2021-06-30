package main

import (
	"fmt"
	"time"
)

func main() {
 	go task()
	time.Sleep(time.Second * 10)
}

func task(){

	for range time.Tick(time.Second *1) {
		fmt.Println("Foo")
		time.Sleep(5 * time.Second)

	}
}