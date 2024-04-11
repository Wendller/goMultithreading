package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Step %d from Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 (main)
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")

	time.Sleep(12 * time.Second)
}
