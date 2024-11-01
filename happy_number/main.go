package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 {
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit 
			n /= 10
		}
		if seen[sum] {
			return false
		}

		seen[sum] = true
		n = sum
	}
	return true 
}

func main() {
	n := 19
	if isHappy(n) {
		fmt.Println(n, "is a happy number")
	} else {
		fmt.Println(n, "is not a happy number")
	}
}