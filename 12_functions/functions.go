package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func getLang() (string, string, string) {
	return "Go", "Golang", "Go Programming Language"
}

func processIt(fn func(a int, b int) int) {
	result := fn(10, 5)
	fmt.Println("Result:", result)
}

func getfunc() func(a int, b int) int {
	return func(a int, b int) int {
		return a / b
	}
}

func main() {
	fmt.Println(add(3, 5))
	fmt.Println(sub(3, 5))
	lang1, lang2, _ := getLang()
	fmt.Println(lang1, lang2)

	fn := func(a int, b int) int {
		return a * b
	}
	processIt(fn)
	processIt(add)

	fn2 := getfunc()
	fmt.Println("Division:", fn2(24, 6))
}
