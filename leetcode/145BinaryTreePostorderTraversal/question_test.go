package binaryTreePostorderTraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run("binarytreeinordertraversal", func(t *testing.T) {
			assert.Equal(t, postorderTraversal(tt.args.root), tt.want)
		})
	}
}
