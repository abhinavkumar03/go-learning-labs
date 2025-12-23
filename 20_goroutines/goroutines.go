package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task: ", id)
}

func main() {
	// launching multiple goroutines
	for i := 0; i < 5; i++ {
		go task(i)
	}

	// anonymous goroutine
	for j := 0; j < 5; j++ {
		go func(i int) {
			fmt.Println("Anonymous function executing: ", i)
		}(j)
	}

	time.Sleep(2 * time.Second)

}
