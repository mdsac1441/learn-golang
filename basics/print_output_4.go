// In Go programming, we use these three functions to print output messages on the screen.

// fmt.Print()
// fmt.Println()
// fmt.Printf()

// Program to illustrate fmt.Print()

package main

// importing the fmt package
import (
	"fmt"
)

func main() {

	fmt.Print("Hello, ")
	fmt.Print("World!")

	name := "John"
	fmt.Print(name)
	fmt.Print("Name: ", name)

	currentSalary := 50000

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println("Current Salary:", currentSalary)
	currentAge := 21
	fmt.Printf("Age = %d", currentAge)

	// 	In Go, every data type has a unique format specifier.

	// 	Data Type	Format Specifier
	// 	integer	%d
	// 	float	%g
	// 	string	%s
	// 	bool	%t

	var annualSalary = 65000.5

	fmt.Printf("Annual Salary: %g", annualSalary)

	var name1 = "John"
	age := 23

	fmt.Printf("%s is %d years old.", name1, age)

	// It's also possible to print output without using the fmt package. For that, we use print() and println().
	println("Using println instead of fmt.Println")

	print("Using print instead of fmt.Print")
}
