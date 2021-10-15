package binarytreeinordertraversal

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
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run("binarytreeinordertraversal", func(t *testing.T) {
			assert.Equal(t, inorderTraversal(tt.args.root), tt.want)
		})
	}
}
