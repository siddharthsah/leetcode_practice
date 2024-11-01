package main

import (
	"fmt"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	dfs(root, &maxLength)
	return maxLength
}

func dfs(t *TreeNode, maxLength *int) int {
	if t == nil {
		return 0
	}

	left := dfs(t.Left, maxLength)
	right := dfs(t.Right, maxLength)
	*maxLength = max(*maxLength, left+right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	start := time.Now()
	// Create a sample binary tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	diameter := diameterOfBinaryTree(root)
	elapsed := time.Since(start)

	fmt.Println("Diameter of the binary tree: ", diameter)
	fmt.Printf("Binomial took %s", elapsed)
}
