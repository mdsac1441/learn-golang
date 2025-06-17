package main

import "fmt"

// // handle panic call
// func division(num1, num2 int) {

// 	// if num2 is 0, program is terminated due to panic
// 	if num2 == 0 {
// 		panic("Cannot divide a number by zero")
// 	} else {
// 		result := num1 / num2
// 		fmt.Println("Result: ", result)
// 	}
// }

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER:", a)
	}

}
func division(num1, num2 int) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
	// defer handlePanic() // we get different result
}
func main() {

	// We use the defer statement to prevent the execution of a function until all other functions execute.
	// defer the execution of Println() function
	// defer fmt.Println("Three")

	// fmt.Println("One")
	// fmt.Println("Two")
	// fmt.Println()

	// When we use multiple defer in a program, the order of execution of the defer statements will be LIFO (Last In First Out).
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	// panic
	// fmt.Println("Help! Something bad is happening.")
	// panic("Ending the program")
	// fmt.Println("Waiting to execute")

	division(4, 2)
	division(8, 0)
	division(8, 8)

}
