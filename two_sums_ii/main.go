package main

import (
	"fmt"
	"time"
)

func two_sums_ii(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left ++
		} else {
			right --
		}
	}
	return []int{}  // no solution found
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	start := time.Now()
	fmt.Println(two_sums_ii(numbers, target))
	elapsed := time.Since(start)
	fmt.Printf("program took %s\n", elapsed)
}