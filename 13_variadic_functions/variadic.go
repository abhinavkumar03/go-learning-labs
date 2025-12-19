package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func print(args ...interface{}) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println(1, 2, 3, "hi") // it is variadic function

	// passing multiple arguments to variadic function
	res := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", res)

	// passing slice to variadic function
	arr := []int{10, 20, 30}
	fmt.Println("Sum from slice:", sum(arr...)) // unpacking slice

	print("Hello", 42, 3.14, true)
}
