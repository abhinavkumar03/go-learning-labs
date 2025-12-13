package main

import "fmt"

// for is the only loop construct in Go
// It can be used in several ways
func main() {
	// Basic for loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // skip the rest of the loop when i is 3
		}
		fmt.Println("Basic for loop iteration:", i)
	}

	// for loop as a while loop
	i := 0
	for i < 5 {
		fmt.Println("While-like for loop iteration:", i)
		i++
	}

	// Infinite loop with break
	for {
		fmt.Println("Infinite loop iteration, will break after first")
		break
	}

	// range loop
	for j := range 10 {
		fmt.Println("Range loop iteration:", j)
	}

}
