package balancedBinaryTree

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Right) && isBalanced(root.Left)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
