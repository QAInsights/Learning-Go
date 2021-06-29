package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func (p person)greet(){
	fmt.Println("Hello", p.name)
}

func main() {


	sum(1,1,1,1,1)
	sayHello("Hello Go!", 24)
	sayMoreHello("Hello", "NaveenKumar", 24)
	a := 5
	b := 5
	twoSum(a, b)
	twoSumMemory(&a, &b)
	fmt.Println("Now the value of c and d", a, b)
	
	multiplyAll("Hello there!", 2,2,2,2,2)
	m := multiplyByPointer(2,2,2,2,2)
	fmt.Println("Pointer multiply", *m)

	div, err := divide(10.0, 3.0)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(div)

	// Anonymous Functions

		func (){
			fmt.Println("From Anonymous func: => Hello there!!!")
		}()

		for z := 0; z < 5; z++ {
			func(j int){
				fmt.Println(j)
			}(z)
		}

	// Func struct
		personA := person{
			name: "NaveenKumar",
			age: 24,
		}
		fmt.Println(personA.name, personA.age)
		personA.greet()
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return b, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
func multiplyByPointer(val ...int)*int{
	ans := 1
	for _, j := range val{
		ans = ans * j
	}
	return &ans
}

func multiplyAll(msg string, values ...int)  {

	ans := 1
	for _, j := range values{
		fmt.Println(j)
		ans = ans * j
	}
	fmt.Println("Multiplication: ", msg, ans)
}

func twoSumMemory(c, d *int){
	*c = 20
	*d = 20
	fmt.Println("Sum of two numbers", *c + *d) // Prints 40 because we are passing address

}

func twoSum(a, b int){
	fmt.Println("Sum of two numbers", a + b) // Prints 10
	a = 24
	fmt.Println("Value of a: ", a) // Prints 24

}

func sayMoreHello(msg, name string, age int){ //Syntatic Sugar
	fmt.Println(msg, name, age)
}
func sayHello(msg string, age int){
	fmt.Println(msg, age)
}
func sum(values ...int){

	ans := 0
	for i:= range values {
		ans = ans + i
	}
	fmt.Println(ans)
}