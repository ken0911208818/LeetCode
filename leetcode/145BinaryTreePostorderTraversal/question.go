package binaryTreePostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func postorderTraversal(root *TreeNode) []int {
    var output []int
	return Postorder(root, &output)
}

func Postorder(root *TreeNode, output *[]int) []int {
	if root != nil {
        Postorder(root.Left, output)
        Postorder(root.Right, output)
        *output = append(*output, root.Val)
	}
	return *output
}
