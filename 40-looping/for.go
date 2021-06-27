package main

import (
	"fmt"
)

func main(){

	//
	//for i := 1; i <= 5; i = i + 2 {
	//	fmt.Println(i)
	//}

	j := 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}
	fmt.Println(j) // prints 10 from line 14

	for i, j := 0,0; i <= 5; i, j = i +1, j + 2 {
		fmt.Println(i, j)
	}

	m := 0

	// Do Loop - kind of
	for m < 5 {
		fmt.Println(m)
		m++
	}

	// While

	n := 0

	for {
		fmt.Println("While Loop (kind of) 🍕 ", n)
		n++
		if n == 10 {
			break
		}
	}
	n = 0
	for ;n<=10;n++ {
		if n%2 == 0 {
			continue
			fmt.Println(n)
		}
		fmt.Println(n)
	}

	// Labeling the Loop
	UpperLoop:
		for a:= 0; a <= 3; a++ {
			for b:= 0; b <= 3; b++ {
				fmt.Println("a * b", a*b)
				if a * b >= 3 {
					break UpperLoop // Breaks the first for loop
				}
			}
		}

	// Slices Looping

	fruits := []string{"apple", "orange", "mango"}

	for k, v := range fruits {
		fmt.Println(k, v)
	}

	// Array Looping

	colors := [3]string{"red", "blue", "green"}

	for k, v := range colors {
		fmt.Println(k,v)
	}

	// Maps Looping

	zipCodes := map[string]int{"OH":45000, "IL": 60003, "CA": 10001}
	for k, v := range zipCodes{
		fmt.Println(k,v)
	}

	for k, _ := range zipCodes{
		fmt.Println(k) // Prints only the States: OH, IL and CA
	}

	// String Looping

	myBlog := "QAInsights"

	for k, v := range myBlog {
		fmt.Println(k, v) // Prints Unicode representation of characters
		fmt.Println(k, string(v)) // Prints the actual character
	}
	// Infinity Loop

	//for {
	//	fmt.Println("Infinity")
	//	time.Sleep(1 * time.Second)
	//}

	// Infinity Loop

	//s := 0
	//for ; s<=1 ; {
	//	fmt.Println("Infinity Loop")
	//	time.Sleep(1 * time.Second)
	//}

}

