package main

import (
	"fmt"
)

func main() {

	var floatValue1 float32 = 5.45

	// type conversion from float to int
	var intValue1 int = int(floatValue1)

	fmt.Printf("Float Value is %g\n", floatValue1)
	fmt.Printf("Integer Value is %d\n", intValue1)

	var intValue2 int = 2

	// type conversion from int to float
	var floatValue2 float32 = float32(intValue2)

	fmt.Printf("Integer Value is %d\n", intValue2)
	fmt.Printf("Float Value is %f", floatValue2)

	// Go doesn't support the implicit type casting.

	var num1 int = 20
	var num2 float32 = 5.7
	var sum float32

	// addition of different data types
	sum = float32(num1) + num2

	fmt.Printf("\nSum is %g\n", sum)
}
