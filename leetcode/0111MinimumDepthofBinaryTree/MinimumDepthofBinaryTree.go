package minimumDepthOfBinaryTree

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left) + 1
	right := depth(root.Right) + 1

	if left == 1 {
		return right
	}
	if right == 1 {
		return left
	}
	if left < right {
		return left
	}
	return right
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return min(depth(root.Left), depth(root.Right)) + 1
}

func min(a int, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}
	if a > b {
		return b
	}
	return a
}
