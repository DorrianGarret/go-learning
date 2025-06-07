package main

import "fmt"

func main() {
	capitals := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
	}
	capitals["Russia"] = "Moscow"
	capitals["France"] = "Kiev"

	delete(capitals, "France")

	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}
}
