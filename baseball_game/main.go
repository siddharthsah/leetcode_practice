package main

import (
	"fmt"
	"strconv"
)

// Time Complexity: O(N), where N is the number of operations. Each operation is processed in constant time.
// Space Complexity: O(N), as the stack can grow to a maximum size of N elements in the worst case.

func baseballGame(ops []string) int {
	stack := []int{}
	for _, val := range ops {
		switch val {
		case "+":
			x, y := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack, x+y)
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "C":
			stack = stack[:len(stack)-1]
		default:
			op, err := strconv.Atoi(val)
            if err != nil {
                // Handle error, e.g., log the error or return an error value
                panic(err) // Or handle the error appropriately
            }
			stack = append(stack, op)
		}
	}
	sum := 0
	for _, val := range stack {
		sum += val
	}
	return sum
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(baseballGame(ops))
}
