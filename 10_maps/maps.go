package main

import (
	"fmt"
	"maps"
)

// maps -> key-value pairs, dynamic size, unordered collection
func main() {
	// declaring a map
	m := make(map[string]string)

	// adding key-value pairs
	m["name"] = "Alice"
	m["city"] = "Wonderland"

	// retrieving value by key
	fmt.Println(m["name"])
	fmt.Println(m["age"]) // key does not exist, returns zero value

	fmt.Println(len(m)) // length of the map
	// deleting a key-value pair
	fmt.Println(m, len(m))
	delete(m, "city")
	fmt.Println(m, len(m))

	// clearing a map
	clear(m)
	fmt.Println(m, len(m))

	marks := map[string]int{
		"Math":    95,
		"Science": 90,
	}
	fmt.Println(marks)

	// checking if a key exists
	_, ok := marks["English"] // comma ok idiom, _ to ignore the value
	if ok {
		fmt.Println("English key exists")
	} else {
		fmt.Println("English key does not exist")
	}

	k, ok := marks["Math"] // checking for a non-existing key and getting its value
	if ok {
		fmt.Println("Math key exists with value", k)
	} else {
		fmt.Println("Math key does not exist")
	}

	// equality of maps
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 2, "a": 1}
	// throwing error: invalid operation: m1 == m2 (map can only be compared to nil)
	// fmt.Println(m1 == m2) // true, maps are equal if they have the same keys and values

	fmt.Println(maps.Equal(m1, m2)) // using maps.Equal to compare maps
}
