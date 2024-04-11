package main

import "fmt"

// channel just receive data
func receive(name string, channel chan<- string) {
	channel <- name
}

// channel just send data
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	messageChan := make(chan string)

	go receive("Hello", messageChan)
	read(messageChan)
}
