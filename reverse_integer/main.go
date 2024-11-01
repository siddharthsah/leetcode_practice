package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var reversed int
	for x != 0 {
		rem := x % 10
		x /= 10
		reversed = reversed * 10 + rem
		if (reversed > math.MaxInt32 && x > 0) || (reversed < math.MaxInt32 && x < 0) {
			return 0
		} 
	}
	return reversed
}

func main () {
	x := 123
	reversed := reverse(x)
	fmt.Println(reversed)
}