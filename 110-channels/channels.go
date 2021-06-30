package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan string, 50)
	wg.Add(2)

	go func(ch <- chan string){ // Pulling out from the channel - receive only
		for i := range ch{ //pulling a single value is the ACTUAL value from the channel - i holds the string
			fmt.Println(i)
		}
		wg.Done()

	}(ch)

	go func(ch chan <- string) { // From string to channel i.e. pushing into the channel - send only
		ch <- "NaveenKumar" // sending data to the channel
		ch <- "Jananya"
		ch <- "Diya"
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
