package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices -> dynamic size, more flexible than arrays
	// most used  constructor in Go
	var nums []int                            // declaring a slice of integers
	fmt.Println(nums, nums == nil, len(nums)) // checking if slice is nil and its length

	var marks = make([]int, 2, 3) // creating a slice with length 5
	marks[0] = 95                 // assigning values to the slice indices
	fmt.Println(marks, cap(marks))

	marks = append(marks, 90, 80, 70)          // appending values to the slice
	fmt.Println(marks, len(marks), cap(marks)) // length and capacity after appending

	var values = make([]int, 0, 2)
	values = append(values, 1)
	values = append(values, 2)
	values = append(values, 3)
	values = append(values, 4)
	values = append(values, 5)
	fmt.Println(values, len(values), cap(values))

	heights := []int{} // slice literal
	heights = append(heights, 170)
	fmt.Println(heights, len(heights), cap(heights))

	// copying slices
	var nums1 = make([]int, 0, 5)
	nums1 = append(nums1, 1)
	var nums2 = make([]int, len(nums1))
	copy(nums2, nums1) // copying contents of nums1 to nums2
	fmt.Println(nums1, nums2)

	// slicing slices
	numbers := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(numbers[1:4])

	// slice equality
	// slices cannot be compared directly using == operator
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(slice1, slice2))

	// multi-dimensional slices
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
}
