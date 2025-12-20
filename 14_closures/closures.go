package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Create a closure that maintains its own state
	increment := counter() // increment holds the closure returned by counter
	fmt.Println(increment())
	fmt.Println(increment())

	increment2 := counter() // increment2 is a new closure with its own state
	fmt.Println(increment2())
	fmt.Println(increment2())
}
