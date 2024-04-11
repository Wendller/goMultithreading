package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string)

	//Thread 2
	go func() {
		channel <- "Hello channel"
	}()

	//Thread 1
	message := <-channel
	fmt.Printf(message)
}
