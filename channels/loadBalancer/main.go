package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for message := range data {
		fmt.Printf("Worker %d received: %d\n", workerId, message)
		time.Sleep(time.Second)
	}
}

func main() {
	dataChannel := make(chan int)

	// 3 Threads
	go worker(1, dataChannel)
	go worker(2, dataChannel)
	go worker(3, dataChannel)

	for i := 0; i < 15; i++ {
		dataChannel <- i
	}
}
