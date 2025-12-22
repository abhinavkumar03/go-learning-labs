package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		println(item)
	}
}

func printGenericSlice[T any](items []T) {
	for _, item := range items {
		println(item)
	}
}

func printIntegerAndStringSlice[T int | string](items []T) {
	for _, item := range items {
		println(item)
	}
}

type Printable[T any] struct {
	value []T
}

func main() {
	// without using generics
	printSlice([]int{1, 2})
	printStringSlice([]string{"golang", "Js"})

	// using generics
	printGenericSlice([]int{1, 2})
	printGenericSlice([]string{"golang", "Js"})
	printGenericSlice([]bool{true, false})

	// using type constraints only int and string
	printIntegerAndStringSlice([]int{1, 2})
	printIntegerAndStringSlice([]string{"golang", "Js"})
	// printIntegerAndStringSlice([]bool{true, false}) // compilation error

	intPrintable := Printable[int]{value: []int{1, 2, 3}}
	fmt.Println(intPrintable)
	stringPrintable := Printable[string]{value: []string{"a", "b", "c"}}
	fmt.Println(stringPrintable)
}
