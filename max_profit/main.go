package main

import (
	"fmt"
	"math"
	"time"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		minPrice = min(minPrice, price)
		profit := price - minPrice
		maxProfit = max(maxProfit, profit)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	start := time.Now()
	fmt.Println(maxProfit(prices))
	elapsed := time.Since(start)
	fmt.Printf("time elapsed since start of the program %s", elapsed)
}
