// O(logn) searching algorithm

package main

import (
	"fmt"
	"math"
)

// sorting algorithm
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// binary search algorithm
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	/* Binary Search Example */
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}

	/* Merge Sort Example */
	n := 16
	fmt.Printf("Logarithm base 2 of %d is %f\n", n, math.Log2(float64(n)))

	unsortedArr := []int{38, 27, 43, 3, 9, 82, 10}
	sortedArr := mergeSort(unsortedArr)
	fmt.Println("Unsorted array:", unsortedArr)
	fmt.Println("Sorted array:", sortedArr)
}