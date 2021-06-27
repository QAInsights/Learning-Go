package main

import (
	"fmt"
)

func main() {

	// short circuiting

	guess := -1

	if guess < 0 || returnTrue() || guess > 100{
		fmt.Println("Must be 1 - 100")
	} else {
		fmt.Println("Valid number")
	}

	i := 1
	switch  {
	case i < 10:
		fmt.Println("Case 1")
		fallthrough
	case i < 20:
		fmt.Println("Case 2")
	default:
		fmt.Println("Case default")
	}

	var a interface{} = 1

	switch a.(type) {
	case int:
		fmt.Println("Int")
		fmt.Println("Second Line")
		break
		fmt.Println("This will not print")
	case bool:
		fmt.Println("Bool")
	case string:
		fmt.Println("String")
	case int32:
		fmt.Println("Int 32")
	default:
		fmt.Println("Nothing above")
	}
}
func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}