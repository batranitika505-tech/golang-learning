// Here we will learn about for loops in Go (Golang)

package main

import "fmt"

func main() {

	// 1. Basic for loop (works like a while loop in other languages)
	i := 1
	for i <= 3 {
		fmt.Println("i =", i)
		i = i + 1
	}

	// 2. Classic for loop (with init, condition, and post statement)
	for j := 0; j < 3; j++ {
		fmt.Println("j =", j)
	}

	// 3. Using range in a for loop (to iterate over collections)
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 4. Using range with only index
	for index := range nums {
		fmt.Println("Only index:", index)
	}

	// 5. Using range with only value (ignore index using _)
	for _, value := range nums {
		fmt.Println("Only value:", value)
	}
}
