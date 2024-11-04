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
// space complexity is O(n) in worst case as n is the height of tree, in a balanced tree, it will be O(log n)
// time complexity is O(n), traversing each node
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	start := time.Now()
	// Create a sample binary tree
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right.Right = &TreeNode{Val: 1}

	targetSum := 22

	fmt.Println(hasPathSum(root, targetSum))
	elapsed := time.Since(start)
	fmt.Printf("program took %s", &elapsed)

}
