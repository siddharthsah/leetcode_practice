package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)

	if n == 2 {
		return min(cost[0], cost[1])
	}

	for i := 2; i <=n; i++ {
		a := dp[i-1] + cost[i-1]
		b := dp[i-2] + cost[i-2]

		dp[i] = min(a, b)
	}
	return dp[n]
}
// func minCostClimbingStairs(cost []int) int {
//     n := len(cost)
//     dp := make([]int, n)

//     // Base cases
//     dp[0] = cost[0]
//     dp[1] = cost[1]

//     for i := 2; i < n; i++ {
//         dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
//     }

//     return min(dp[n-1], dp[n-2])
// }

func min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func main() {
	cost := []int {10, 15, 20}
	minCost := minCostClimbingStairs(cost)
	fmt.Println("minimum cost to climb staris: ", minCost)
}