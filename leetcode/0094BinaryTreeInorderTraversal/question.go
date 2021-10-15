package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var output []int
	return inorder(root, &output)
}

func inorder(root *TreeNode, output *[]int) []int {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
	return *output
}
