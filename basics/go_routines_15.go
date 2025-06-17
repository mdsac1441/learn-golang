// Running two goroutines concurrently

package main

import (
	"fmt"
	"time"
)

// creating a function
func display(message string) {

	fmt.Println(message)
	// // infinite for loop
	// for {
	// 	fmt.Println(message)

	// 	// to sleep the process for 1 second
	// 	time.Sleep(time.Second * 1)
	// }
}

func main() {

	// // Calling function with goroutine
	// go display("Process 1")

	// display("Process 2")
	// run two different goroutine
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	// to sleep main goroutine for 1 sec
	time.Sleep(time.Second * 1)
}
