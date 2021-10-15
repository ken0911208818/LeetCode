package binaryTreePreorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func preorderTraversal(root *TreeNode) []int {
    var output []int
	return Preorder(root, &output)
}

func Preorder(root *TreeNode, output *[]int) []int {
	if root != nil {
        *output = append(*output, root.Val)
        Preorder(root.Left, output)
        Preorder(root.Right, output)
	}
	return *output
}
