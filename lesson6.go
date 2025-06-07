package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello, ", name)
}

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	greet("Bob")
	fmt.Println(add(1, 2))
	fmt.Println(divide(1, 2))
}
