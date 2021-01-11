package selectexample

import (
	"fmt"
	"time"
)

// This is good review: https://tour.golang.org/concurrency/5
func simpleSelect() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(2)
		close(c)
	}()

	select {
	case <-c:
		fmt.Printf("done: %v\n", time.Since(start))
	}
}
