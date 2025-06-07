package main

import "fmt"

func main() {
	var city [4]string = [4]string{
		"Ukraine",
		"New York",
		"Singapore",
		"London",
	}
	fmt.Println(city[0])
	fmt.Println(city[1])
	fmt.Println(city[2])
	fmt.Println(city[3])

	numbers := []int{1, 2, 3}

	numbers = append(numbers, 4, 5)

	sub := numbers[1:4]
	fmt.Println(sub)
}
