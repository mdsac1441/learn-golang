package main

import "fmt"

func main() {

	// create two channels
	number := make(chan int)
	message := make(chan string)

	// function call with goroutine
	go channelNumber(number)
	go channelMessage(message)

	// selects and executes a channel
	select {

	case firstChannel := <-number:
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data:", secondChannel)
	}

}

// goroutine to send integer data to channel
func channelNumber(number chan int) {
	number <- 15
}

// goroutine to send string data to channel
func channelMessage(message chan string) {
	message <- "Learning Go select"
}
