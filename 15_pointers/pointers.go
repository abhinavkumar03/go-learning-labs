package main

import "fmt"

func changeValue(val *int) {
	fmt.Println("memory address inside function:", val)
	*val += 20
}

func main() {
	var x int = 10
	fmt.Println("Before:", x)
	fmt.Println("memory address of x:", &x)
	changeValue(&x)
	fmt.Println("After:", x)

	// changeValue(x) // This will cause a compile-time error
}
