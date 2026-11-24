package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Search returns the node holding key, or nil. The BST ordering makes the
// path deterministic, so a loop beats recursion: no stack, no frames.
func Search(root *TreeNode, key int) *TreeNode {
	for root != nil && root.Val != key {
		if key < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(8, tn(3, tn(1, nil, nil), tn(6, nil, nil)), tn(10, nil, tn(14, nil, nil)))
	fmt.Println(Search(root, 6).Val, Search(root, 7))
}
