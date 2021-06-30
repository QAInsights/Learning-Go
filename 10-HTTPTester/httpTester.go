package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

func main() {
	fmt.Println("HTTP Tester")

	url := "https://example.com"
	var vus = 10
	const duration = time.Second * 10 // seconds
	for i := 0; i < vus; i++ {
		wg.Add(1)
		ctx := context.Background()
		go loadTester(ctx, url, i, duration)
	}
	wg.Wait()
}

func loadTester(ctx context.Context, url string, vus int, duration time.Duration) {

	select {
		case <-time.After(duration):
			fmt.Println("after")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-:
			fmt.Println("during")
	}
	fmt.Println("vu started", vus+1, time.Now())

	d := time.Now().Add(time.Second * duration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}


	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response received", time.Since(start).Milliseconds(),"ms")

	wg.Done()
}