package main

import "fmt"

const age = 25 // untyped constant
// name := "Alice" // This will cause a compile-time error

// Grouped constant declaration
const (
	port    = 5000
	host    = "localhost"
	timeout = 30 // seconds
)

func main() {
	const Pi = 3.14               // untyped constant
	const Greeting string = "Hi!" // typed constant

	// Pi = 3.14159 // This will cause a compile-time error
	fmt.Println("Age:", age)
	fmt.Println("Pi:", Pi)
	fmt.Println("Greeting:", Greeting)
}
