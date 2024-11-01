package main

import "fmt"

func numOnes(x int) int {
	count := 0
	// this operation clears the rightmost bit and increase the count
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	num := 21
	fmt.Println("number of 1 bits are", numOnes(num))
}
