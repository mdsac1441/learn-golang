// Program to create an array and prints its elements

package main

import (
	"fmt"
	"reflect"
)

// reflect: used for checking type of data

func main() {

	// declaring array variable of type integer
	// defined size [size=5]
	var arrayOfInteger = [5]int{1, 5, 8, 0, 3}

	fmt.Println(arrayOfInteger[1:5])

	// declaring array variable of type string
	// undefined size
	var arrayOfString = [...]string{"Hello", "Ahmed"}
	fmt.Println(arrayOfString)
	fmt.Println("Length:", len(arrayOfString))
	fmt.Println("Capacity:", cap(arrayOfString))
	fmt.Println("Type:", reflect.TypeOf(arrayOfString))

	// Note: Here, if [] is left empty, it becomes a slice. So [...] is a must if we want an undefined size array.

	languages := [3]string{"Go", "Java", "C++"}

	// access element at index 0
	fmt.Println(languages[0]) // Go

	// access element at index 2
	fmt.Println(languages[2]) // C++

	// declaring an array
	var arrayOfIntegers [3]int

	// elements are assigned using index
	arrayOfIntegers[0] = 5
	arrayOfIntegers[1] = 10
	arrayOfIntegers[2] = 15

	fmt.Println(arrayOfIntegers)

	// to initialize the elements of index 0 and 3 only
	arrayOfIntegers1 := [5]int{0: 7, 3: 9}

	fmt.Println(arrayOfIntegers1)

	weather := [3]string{"Rainy", "Sunny", "Cloudy"}

	// change the element of index 2
	weather[2] = "Stromy"

	fmt.Println(weather)

	age := [...]int{12, 4, 5}

	// loop through the array
	for i := 0; i < len(age); i++ {
		fmt.Println(age[i])
	}

	fmt.Println()
	// creating a 2 dimensional array
	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	// accessing the values of 2d array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}

}
