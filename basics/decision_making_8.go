// Program to display a number if it is positive

package main

import "fmt"

func main() {
	number := 5

	// true if number is less than 0
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	}

	fmt.Println("Out of the loop")

	// checks if number is greater than 0
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	} else {
		fmt.Printf("%d is a negative number\n", number)
	}

	number1 := 12
	number2 := 20

	if number1 == number2 {
		fmt.Printf("Result: %d == %d\n", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d\n", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d\n", number1, number2)
	}

	// outer if statement
	if number1 >= number2 {

		// inner if statement
		if number1 == number2 {
			fmt.Printf("Result: %d == %d", number1, number2)
			// inner else statement
		} else {
			fmt.Printf("Result: %d > %d", number1, number2)
		}

		// outer else statement
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}
