package main

import (
	"fmt"
	"sync"
)

func task(id int) {
	fmt.Println("Doing task: ", id)
}

func wgtask(id int, w *sync.WaitGroup) {
	defer w.Done() // will run at last of function execution
	fmt.Println("wait task: ", id)
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

	// time.Sleep(2 * time.Second)

	var wg sync.WaitGroup

	// wait for goroutines to finish
	for k := 0; k < 5; k++ {
		wg.Add(1)
		go wgtask(k, &wg)
	}

	wg.Wait()

}
