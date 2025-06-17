package main

// import the errors package
import (
	"errors"
	"fmt"
)

// function that checks if name is Programiz
func checkName(name string) error {

	// create a new error
	newError := errors.New("Invalid Name")

	// return the error if name is not Programiz
	if name != "Program" {
		return newError
	}

	// return nil if there is no error
	return nil
}

func main() {

	name := "Hello"

	// call the function
	err := checkName(name)

	// check if the err is nil or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Name")
	}

}
