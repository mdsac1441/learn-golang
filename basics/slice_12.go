// Program to create a slice and print its elements

package main

import "fmt"

// this is an array
// numbers := [5]int{1, 2, 3, 4, 5}

// this is a slice
// numbers := []int{1, 2, 3, 4, 5}
func main() {

	// declaring slice variable of type integer
	numbers := []int{1, 2, 3, 4, 5}

	// print the slice
	fmt.Println("Numbers: ", numbers)

	// an integer array
	numbers1 := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	// creating slice from an array
	sliceNumbers := numbers1[4:7]

	fmt.Println(sliceNumbers)

	primeNumbers := []int{2, 3}

	// add elements 5, 7 to the slice
	primeNumbers = append(primeNumbers, 5, 7)
	fmt.Println("Prime Numbers:", primeNumbers)

	// create two slices
	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("Numbers append:", evenNumbers)

	copy(numbers, primeNumbers)

	// print numbers
	fmt.Println("Numbers:", numbers)

	// access elements from index 0 to index 2
	fmt.Println(numbers[0:3])

	// create a slice using make()
	num := make([]int, 5, 7)

	// add elements to numbers
	num[0] = 13
	num[1] = 23
	num[2] = 33
	num[3] = 1
	num[4] = 8

	fmt.Println("Numbers:", num)
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num))
	// for range loop that iterates through a slice
	for _, value := range numbers {
		fmt.Println(value)
	}
}

// Slice Functions
// append()	adds element to a slice
// copy()	copy elements of one slice to another
// Equals()	compares two slices
// len()	find the length of a slice
