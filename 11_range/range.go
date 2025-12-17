package main

import "fmt"

func main() {
	// range -> iterate over elements in arrays, slices, maps, strings

	// iterating over an array
	arr := [5]int{10, 20, 30, 40, 50}

	// using range to get index and value
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// using range to iterate over keys
	for ele := range m {
		fmt.Println("Element:", ele)
	}

	// using range to get key and value
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// using range to iterate over string characters
	// i is starting byte index of each rune in the string
	for i, c := range "hello" {
		fmt.Printf("Index: %d, Character: %c\n", i, c)
		fmt.Println(i, c, string(c)) // i is the byte index, c is the rune (Unicode code point)
	}
}
