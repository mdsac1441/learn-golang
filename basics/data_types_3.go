package main

import "fmt"

func main() {
	var integer1 int
	var integer2 int
	integer1 = 5
	integer2 = 10

	fmt.Println(integer1)
	fmt.Println(integer2)

	var salary1 float32
	var salary2 float64

	salary1 = 50000.503882901

	// can store decimals with greater precision
	salary2 = 50000.503882901

	fmt.Println(salary1)
	fmt.Println(salary2)

	var message string
	message = "Welcome to Golang World"

	fmt.Println(message)

	var boolValue bool
	boolValue = false

	fmt.Println(boolValue)
}
