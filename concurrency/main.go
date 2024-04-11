package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var accessNumber uint64 = 0

func main() {
	// mutex := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// mutex.Lock()
		atomic.AddUint64(&accessNumber, 1)
		//mutex.Unlock()

		time.Sleep(300 * time.Millisecond)

		w.Write([]byte(fmt.Sprintf("Your access ordered number is: %d", accessNumber)))
	})

	http.ListenAndServe(":8080", nil)
}

// Try to run with Apache Bench => $ ab -n 10000 -c 100 http://localhost:8080/
