package convertSortedArraytoBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{Val: nums[len(nums)/2], Left: sortedArrayToBST(nums[:len(nums)/2]), Right: sortedArrayToBST(nums[len(nums)/2+1:])}
}

func Preorder(root *TreeNode, output *[]int) []int {
	if root != nil {
        *output = append(*output, root.Val)
        Preorder(root.Left, output)
        Preorder(root.Right, output)
	}
	return *output
}