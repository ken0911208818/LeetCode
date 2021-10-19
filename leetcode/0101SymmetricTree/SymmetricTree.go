package symmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSame(root.Left, root.Right)
}

func isSame(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil{
		return false
	}
	if !isSame(l.Left,r.Right) {
		return false
	}

	if !isSame(l.Right, r.Left) {
		return false
	}
	return l.Val == r.Val
}
