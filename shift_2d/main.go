package main

import "fmt"
/*
Handle Large Shifts:
The k %= m * n operation ensures that we only need to shift the grid by the remainder of k divided by the total number of elements. This prevents unnecessary iterations.
Shifting Elements:
The outer loop iterates k times, simulating k right shifts.
In each iteration, we:
Store the bottom-right element in a prev variable.
Shift all elements to the right, overwriting the previous values.
Assign the stored prev element to the top-left corner.

Time Complexity: O(m * n * k)
The outer loop iterates k times.
The inner nested loops iterate over all m * n elements of the grid.
Space Complexity: O(1)
The algorithm modifies the input grid in-place, without using additional space proportional to the input size.
*/
func shift2D(grid [][]int, k int) [][]int {
	m, n  := len(grid), len(grid[0])  //number of rows and columns
	k %= m * n // Handle cases where k is larger than total number of elements in grid

	for i := 0; i < k; i++ {
		prev := grid[m - 1][n - 1]
		for row := 0; row < m; row++ {
			for col := 0; col < n; col++ {
				next := grid[row][col]
				grid[row][col] = prev
				prev = next
			}
		}
	}
	return grid
}

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	fmt.Println(shift2D(grid, k))
}