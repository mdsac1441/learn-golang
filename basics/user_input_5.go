package main

import "fmt"

func main() {
	var name string
	var age uint

	// takes input value for name
	fmt.Print("Enter your name: ")
	fmt.Scan(&name, &age)

	fmt.Printf("Name: %s\nage: %d\n", name, age)

	// take name and age input
	fmt.Println("Enter your name and age:")
	fmt.Scanln(&name, &age)

	fmt.Printf("Name: %s\nage: %d\n", name, age)

	fmt.Println("Enter your Full Name:")
	fmt.Scanln(&name)
	fmt.Printf("Name: %s\n", name)

	// take name and age input using format specifier
	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Name: %s\nAge: %d", name, age)

	var temperature float32
	var sunny bool

	// take float input
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%g", &temperature)

	// take boolean input
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)

}
