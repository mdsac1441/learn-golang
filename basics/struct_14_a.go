package main

import (
	"fmt"
)

func main() {

	// declaring a struct
	type Person struct {
		name string
		age  int
	}

	// assign value to struct while creating an instance
	person1 := Person{"John", 25}
	fmt.Println("Person-1:", person1)

	// define an instance
	var person2 Person

	// assign value to struct variables
	person2 = Person{
		name: "Sara",
		age:  29,
	}

	fmt.Println("Person-2:", person2)

	// declaring a struct
	type Rectangle struct {
		length  int
		breadth int
	}

	// declaring instance rect1 and defining the struct
	rect := Rectangle{22, 12}

	// access the length of the struct
	fmt.Println("Length:", rect.length)

	// access the breadth of the struct
	fmt.Println("Breadth:", rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area:", area)
}
