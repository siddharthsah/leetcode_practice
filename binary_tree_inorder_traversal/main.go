package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//this is a recursive method

// Time Complexity: O(N), where N is the number of nodes in the tree. This is because we visit each node exactly once.
// Space Complexity: O(H), where H is the height of the tree. This is due to the recursion stack, which can grow up to 
// the height of the tree in the worst case. For balanced binary trees, the height is log(N), so the space complexity can be considered O(log N).

func inOrderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return result

}

//this is a iterative method
// func inOrderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	stack := []*TreeNode{}
// 	cur := root

// 	for cur != nil || len(stack) > 0 {
// 		if cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		} else {
// 			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
// 			res = append(res, cur.Val)
// 			cur = cur.Right
// 		}
// 	}
// 	return res
// }
func main() {
	// Create a sample binary tree
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	result := inOrderTraversal(root)
	fmt.Println(result) // Output: [1 3 2]
}
