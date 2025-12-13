package main

import "fmt"

func main() {
	age := 18

	// Simple if-else statement
	if age < 18 {
		fmt.Println("Minor")
	} else if age <= 18 && age < 65 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	var role = "admin"
	var hasPersmission = true

	// Using logical operators
	if role == "admin" || hasPersmission {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

	// If with initialization statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	}

	// go does not have ternary operator, but we can use if-else to achieve similar functionality
	// number := 10
	// number == 10 ? fmt.Println("Ten") : fmt.Println("Not Ten") // This line will cause a compilation error
}
