package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func Tree2str(root *TreeNode) string {
// 	if root == nil {
// 		return ""
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return fmt.Sprintf("%d", root.Val)
// 	}
// 	leftStr := Tree2str(root.Left)
// 	rightStr := Tree2str(root.Right)

// 	if root.Left == nil {
// 		return fmt.Sprintf("%d(%s)", root.Val, rightStr)
// 	} else if root.Right == nil {
// 		return fmt.Sprintf("%d(%s)", root.Val, leftStr)
// 	} else {
// 		return fmt.Sprintf("%d(%s)(%s)", root.Val, leftStr, rightStr)
// 	}

// }

// Time Complexity: O(N), where N is the number of nodes in the tree.
// Space Complexity: O(H), where H is the height of the tree, due to the recursion stack.

func Tree2str(root *TreeNode) string {

	var result []string
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, "(")
		result = append(result, strconv.Itoa(root.Val))
		if root.Left == nil && root.Right != nil {
			result = append(result, "()")
		}
		preorder(root.Left)
		preorder(root.Right)
		result = append(result, ")")
	}

	preorder(root)
	return strings.Join(result[1 : len(result)-1], "")

}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}

	str := Tree2str(root)
	fmt.Println(str) // Output: 1(2)(3(4))
}
