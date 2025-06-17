package main

import "fmt"

// We can create a function that returns an anonymous function. For example,
func greet() func() {

	return func() {
		fmt.Println("Hi John")
	}

}

func main() {

	g1 := greet()
	g1()
}
