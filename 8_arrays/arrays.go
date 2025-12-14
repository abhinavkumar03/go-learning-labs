package main

import "fmt"

func main() {
	// arrays
	var num [5]int // declaring an array of integers with size 5 defaults to zero values

	// assigning values
	num[0] = 10
	num[1] = 20
	num[2] = 30
	num[3] = 40
	num[4] = 50

	// accessing values
	fmt.Println("First element:", num[0])

	// iterating over an array
	for i := 0; i < len(num); i++ {
		fmt.Printf("Element at index %d: %d\n", i, num[i])
	}

	// printing the entire array
	fmt.Println("array:", num)

	// length of an array
	fmt.Println(len(num))

	var vals [3]bool
	fmt.Println(vals)

	var name [3]string
	fmt.Println(name)

	// array literal
	nums := [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	var matrix [2][3]int // 2D array
	fmt.Println(matrix)

	// fixed size array, memory optimization, performance benefits, constant time access
}
