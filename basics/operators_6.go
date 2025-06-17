package main

import (
	"fmt"
)

func main() {

	num1 := 6
	num2 := 2

	// + adds two variables
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)

	// - subtract two variables
	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, difference)

	// * multiply two variables
	product := num1 * num2
	fmt.Printf("%d * %d is %d\n", num1, num2, product)

	// * divide two integer variables
	quotient := num1 / num2
	fmt.Printf("quotient: %d / %d = %d\n", num1, num2, quotient)
	// If we want the actual result we should always use the / operator with floating point numbers
	numa := 6.0
	numb := 3.0

	result := numa / numb
	fmt.Printf("quotient: %g / %g = %g\n", numa, numb, result)

	// % modulo-divides two variables
	reminder := num1 % num2
	fmt.Println(reminder)

	// increment of num by 1
	num1++
	fmt.Println(num1) // 6

	// decrement of num by 1
	num1--
	fmt.Println(num1) // 4

}
