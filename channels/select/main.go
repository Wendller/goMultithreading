package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID    int64
	Value string
}

func main() {
	channel1 := make(chan Message)
	channel2 := make(chan Message)
	var messageID int64 = 0

	// RabbitMQ
	go func() {
		for {
			time.Sleep(time.Second * 1)

			atomic.AddInt64(&messageID, 1)
			message := Message{ID: messageID, Value: "Event from RabbitMQ\n"}
			channel1 <- message
		}
	}()

	// Kafka
	go func() {
		for {
			time.Sleep(time.Second * 2)

			atomic.AddInt64(&messageID, 1)
			message := Message{ID: messageID, Value: "Event from Kafka\n"}
			channel2 <- message
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Printf("Message %d: %s\n", message1.ID, message1.Value)

		case message2 := <-channel2:
			fmt.Printf("Message %d: %s\n", message2.ID, message2.Value)

		case <-time.After(time.Second * 3):
			println("Timeout")

			// default:
			// 	println("Default")
		}
	}
}
