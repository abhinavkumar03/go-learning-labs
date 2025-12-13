package main

import "fmt"

func main() {
	var a int = 10 // explicit type declaration
	var b = 20     // implicit type declaration
	c := 30        // short variable declaration

	fmt.Println(a, b, c)

	var name string // default is empty string
	name = "Gopher" // assignment

	fmt.Println("Name:", name)

}
