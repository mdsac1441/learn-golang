package main

import "fmt"

func main() {

	// explicitly declare the data type
	var number1 int = 10
	fmt.Println(number1)

	// assign a value without declaring the data type
	var number2 = 20
	fmt.Println(number2)

	// shorthand notation to define variable
	number3 := 30
	fmt.Println(number3)

	// Error: count variable doesn't have a data type
	// var count

	// count1 is of the int type
	// var count1 int

	// count2 is of the int type
	// var count2 = 10

	// var name, age = "Palistha", 22
	// name, age := "Palistha", 22

	const lightSpeed = 299792458 // initial value

	// Error! Constants cannot be changed
	// lightSpeed = 299792460

	// By the way, we cannot use the shorthand notation := to create constants.
	// Error code
	// const lightSpeed := 299792458
}
