// Program to print the first 5 natural numbers

package main

import "fmt"

func main() {

	// for loop terminates when i becomes 6
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	var n, sum = 10, 0

	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
	}

	fmt.Println("sum =", sum)

	// create an array
	numbers := [5]int{11, 22, 33, 44, 55}

	// for loop with range
	for i := range numbers {
		fmt.Println(numbers[i])
	}
	println()
	for _, value := range numbers {
		// throws an error if index is not used but ignore with the help of underscore
		fmt.Println(value)
	}

	// while loop in Golang
	number := 1

	for number <= 5 {
		fmt.Println(number)
		number++
	}

	multiplier := 1

	// run while loop for 10 times
	for multiplier <= 10 {

		// find the product
		product := 5 * multiplier

		// print the multiplication table in format 5 * 1 = 5
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

	// do while loop injh Golang
	number1 := 1
	// loop that runs infinitely
	for {

		// condition to terminate the loop
		if number1 > 5 {
			break
		}

		fmt.Printf("%d\n", number1)
		number1++

	}

	// using range to iterate over the elements of array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}

	strings := "Golang"
	fmt.Println("Index: Character")

	// i access index of each character
	// item access each character
	for i, item := range strings {
		fmt.Printf("%d= %c \n", i, item)
	}

	// create a map
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks obtained:")

	// using for range to iterate through the key-value pair
	for subject, marks := range subjectMarks {
		fmt.Println(subject, ":", marks)
	}
}
