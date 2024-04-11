package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)

	go publish(channel)
	go reader(channel, &waitGroup)

	waitGroup.Wait()
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Publishing Message: %d\n", i)
		channel <- i
	}

	close(channel)
}

func reader(channel chan int, wg *sync.WaitGroup) {
	for x := range channel {
		fmt.Printf("Message Received: %d\n", x)
		wg.Done()
	}
}
