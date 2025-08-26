// write a brute force algorithm
package main

import (
	"fmt"
)

func bruteForceSearch(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Brute Force Algorithm")

	// Example: Finding a number in an array
	numbers := []int{3, 5, 2, 8, 6, 1, 4, 7}
	target := 6
	found := false

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("Number %d found in the array.\n", target)
	} else {
		fmt.Printf("Number %d not found in the array.\n", target)
	}
}