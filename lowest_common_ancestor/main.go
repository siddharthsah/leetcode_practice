package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestorBST(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// if p and q are in different subtrees, the current nodes is the LCA
	if p.Val < root.Val && q.Val > root.Val || p.Val > root.Val && q.Val < root.Val {
		return root
	}

	//Recursively search in the appropriate subtree
	if p.Val < root.Val {
		return lowestCommonAncestorBST(root.Left, p , q)
	} else {
		return lowestCommonAncestorBST(root.Right, p , q)
	}

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

func main () {
	// Create a sample binary search tree
    root := &TreeNode{Val: 6}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 8}
    root.Left.Left = &TreeNode{Val: 0}
    root.Left.Right = &TreeNode{Val: 4}
    root.Right.Left = &TreeNode{Val: 7}
    root.Right.Right = &TreeNode{Val: 9}

	//Find the LCA for the two nodes
	p := root.Left.Right
	q := root.Right.Left
	lca := lowestCommonAncestorBST(root, p, q)
	fmt.Println("lowest common ancestor:", lca)
	printBinaryTree(lca)
}