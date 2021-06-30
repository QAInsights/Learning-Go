package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 10 * time.Second

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	go msg()
	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println(time.Now())
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())


		go msg()
	}

}
func msg() {
	fmt.Println(time.Now())
}