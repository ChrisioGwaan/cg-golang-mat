// O(1)

package main

import (
	"fmt"
)

func main() {
	// 1. Accessing an array element by index (O(1))
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println("Access element at index 2:", arr[2]) // O(1)

	// 2. Updating an array element (O(1))
	arr[2] = 99
	fmt.Println("Updated array:", arr)

	// 3. Appending to a slice (Amortized O(1))
	arr = append(arr, 60)
	fmt.Println("Appended slice:", arr)

	// 4. Accessing a map value by key (O(1) average case, hash lookup)
	m := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("Get value for key 'Bob':", m["Bob"])

	// 5. Updating a map value by key (O(1))
	m["Alice"] = 26
	fmt.Println("Updated map:", m)

	// 6. Inserting into a map (O(1))
	m["Charlie"] = 35
	fmt.Println("Inserted new key:", m)

	// 7. Checking if a key exists in a map (O(1))
	if val, exists := m["Bob"]; exists {
		fmt.Println("Key 'Bob' exists with value:", val)
	}

	// 8. Pushing onto a stack (slice append, O(1) amortized)
	stack := []int{}
	stack = append(stack, 100) // push
	stack = append(stack, 200)
	fmt.Println("Stack after push:", stack)

	// 9. Popping from a stack (O(1))
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("Popped:", top, "Remaining stack:", stack)

	// 10. Queue enqueue (O(1) amortized using slice append)
	queue := []int{}
	queue = append(queue, 1) // enqueue
	queue = append(queue, 2)
	fmt.Println("Queue after enqueue:", queue)

	// 11. Queue dequeue (O(1) by slicing)
	first := queue[0]
	queue = queue[1:]
	fmt.Println("Dequeued:", first, "Remaining queue:", queue)

	// 12. Setting a variable (primitive assignment, O(1))
	x := 42
	x = 100
	fmt.Println("Variable assignment:", x)

	// 13. Returning a value from a function (O(1))
	fmt.Println("Return constant time value:", getConstant())

	// 14. Swapping two variables (O(1))
	a, b := 5, 10
	a, b = b, a
	fmt.Println("Swapped values: a =", a, ", b =", b)

	// 15. Bitwise operations (O(1))
	n := 8 // binary 1000
	fmt.Println("n << 1:", n<<1) // left shift (O(1))
	fmt.Println("n >> 1:", n>>1) // right shift (O(1))
	fmt.Println("n & 1:", n&1)   // AND (O(1))
	fmt.Println("n | 1:", n|1)   // OR (O(1))
	fmt.Println("n ^ 1:", n^1)   // XOR (O(1))
}

func getConstant() int {
	return 7 // O(1)
}