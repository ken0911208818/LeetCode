package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}

	if q == nil {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return p.Val == q.Val
}
