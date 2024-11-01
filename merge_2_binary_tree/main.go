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

func merge2BinaryTree(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	//Create new node with sum of values
	newNode := &TreeNode{Val: root1.Val + root2.Val}

	//recursively merge subtrees
	newNode.Left = merge2BinaryTree(root1.Left, root2.Left)
	newNode.Right = merge2BinaryTree(root1.Right, root2.Right)

	return newNode
}

func printBinaryTree(root *TreeNode) {
    if root == nil {
        return

    }

    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }

        // Print indentation
        for i := 0; i < level; i++ {
            fmt.Print("  ")
        }

        // Print the node value
        fmt.Println(node.Val)

        // Recursively print left and right subtrees
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }

    dfs(root, 0)
}


func main() {
	// Create two sample binary trees
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}

	start := time.Now()
	// Merge the two trees
	mergedRoot := merge2BinaryTree(root1, root2)
	printBinaryTree(mergedRoot)
	elapsed := time.Since(start)
	fmt.Printf("the program tooks %s", elapsed)

}
