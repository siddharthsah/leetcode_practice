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

func invertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//swap left and right children
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	//Recursively invert left and right subtrees
	root.Left = invertBinaryTree(root.Left)
	root.Right = invertBinaryTree(root.Right)

	return root
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
	// Create a sample binary tree
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	start := time.Now()
	printBinaryTree(root)
	invertedRoot := invertBinaryTree(root)
	elapsed := time.Since(start)
	printBinaryTree(invertedRoot)
	fmt.Println("Inverted binary tree: ", invertedRoot)
	fmt.Printf("Binomial took %s", elapsed)
}
