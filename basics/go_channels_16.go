package main

import (
	"fmt"
)

func main() {

	// // create channel
	// number := make(chan int)
	// message := make(chan string)

	// // function call with goroutine
	// go channelData(number, message)

	// // retrieve channel data
	// fmt.Println("Channel Data:", <-number)
	// fmt.Println("Channel Data:", <-message)

	ch := make(chan string)
	// function call with goroutine
	go sendData(ch)

	// receive channel data
	fmt.Println(<-ch)

}

func channelData(number chan int, message chan string) {

	// send data into channel
	number <- 15
	message <- "Learning Go channel"

}

func sendData(ch chan string) {

	// data sent to the channel
	ch <- "Received. Send Operation Successful"
	fmt.Println("No receiver! Send Operation Blocked")

}
