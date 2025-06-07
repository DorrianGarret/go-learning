package main

import "fmt"

func main() {
	var age int
	var language string

	if age < 14 {
		fmt.Println("Child")
	} else if age >= 14 && age <= 17 {
		fmt.Println("Teenager")
	} else if age >= 18 {
		fmt.Println("Adult")
	}

	switch language {
	case "Go":
		fmt.Println("Go is awesome")
	case "Python":
		fmt.Println("Python is cool")
	default:
		fmt.Println("Other language")
	}
}
