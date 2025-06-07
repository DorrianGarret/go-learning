package main

import "fmt"

func main() {
	user := []string{
		"Alex",
		"Bob",
		"Alice",
		"Ted",
		"Klara",
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	for idx, name := range user {
		fmt.Printf("%d: %s\n", idx+1, name)
	}

}
