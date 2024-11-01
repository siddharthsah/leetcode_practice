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

func isSubTree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubTree(root.Left, subRoot) || isSubTree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val{
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
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
	root1.Left = &TreeNode{Val: 4}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 7}

	root2 := &TreeNode{Val: 4}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}

	start := time.Now()
	// Merge the two trees
	printBinaryTree(root1)
	printBinaryTree(root2)
	fmt.Println("the trees are subtree", isSubTree(root1, root2))
	elapsed := time.Since(start)
	fmt.Printf("the program tooks %s", elapsed)
}
