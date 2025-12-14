package main

import (
	"fmt"
	"time"
)

func main() {
	// switch with proper formatting
	i := 3
	switch i {
	case 0:
		fmt.Println("zero")

	case 1:
		fmt.Println("one")

	default:
		fmt.Println("many")
	}

	// multiple cases on one line
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("x is an int: %d\n", v)
		case string:
			fmt.Printf("x is a string: %s\n", v)
		default:
			fmt.Printf("x is of a different type: %T\n", v)
		}
	}

	whoAmI(42)
	whoAmI("hello")
	whoAmI(3.14)

}
