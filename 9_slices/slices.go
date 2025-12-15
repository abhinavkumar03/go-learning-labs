package main

import "fmt"

func main() {
	// slices -> dynamic size, more flexible than arrays
	// most used  constructor in Go
	var nums []int                            // declaring a slice of integers
	fmt.Println(nums, nums == nil, len(nums)) // checking if slice is nil and its length

	var marks = make([]int, 2, 3) // creating a slice with length 5
	fmt.Println(marks, cap(marks))

	marks = append(marks, 90, 80, 70)          // appending values to the slice
	fmt.Println(marks, len(marks), cap(marks)) // length and capacity after appending

}
