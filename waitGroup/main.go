package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Step %d from Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)

		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(20)

	go task("A", &waitGroup)
	go task("B", &waitGroup)

	waitGroup.Wait()
}
