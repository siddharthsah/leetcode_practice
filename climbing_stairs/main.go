package main

import (
	"fmt"
	"time"
)

func climbing_stairs(count int) int {
	if count <= 1 {
		return count
	}

	a, b := 1, 1 //where a and b are two variable instead of creating a dp of size n+1 which shows the numbers of ways you can climb a stair from 0 and 1 steps
	for i := 2; i <= count; i++ {
		a, b = a+b, a
	}
	return a
}

func main() {
	count := 5
	start := time.Now()
	fmt.Println(climbing_stairs(count))
	elapsed := time.Since(start)
	fmt.Printf("program took %s\n", elapsed)
}
