package main

import (
	"fmt"
)

func uglyNumber(n int) bool {
	if n <= 0 {
		return false
	}

	primes := []int{2, 3, 5}
	for _, p := range primes {
		for n%p == 0 {
			n /= p
		}
	}
	return n == 1
}



func main() {
	nums := 20
	fmt.Println("nums is divisible by 2, 3 and 5", uglyNumber(nums))
}
