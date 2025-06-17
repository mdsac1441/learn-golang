package main

import "fmt"

// outer function
func greet() func() string {

	// variable defined outside the inner function
	name := "Ahmed Chistey"

	// return a nested anonymous function
	return func() string {
		name = "Hi " + name
		return name
	}
}

func main() {

	// call the outer function
	message := greet()

	// call the inner function
	fmt.Println(message())
	// fmt.Println(greet()()) // this will also work fine
	fmt.Println(oddNumber(8)())

}

// The returned function is now assigned to the message variable.

// At this point, the execution of the outer function is completed, so the name variable should be destroyed. However, when we call the anonymous function using

func oddNumber(n uint) func() string {
	return func() string {
		if n%2 != 0 {
			return "Odd"
		} else {
			return "Even"
		}
	}
}
