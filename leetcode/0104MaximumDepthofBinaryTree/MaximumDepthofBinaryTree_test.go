package maximunDepthofBinaryTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				root: &TreeNode{
					Val:   3,
					Left:  &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("MaximunDepthofBinaryTree", func(t *testing.T) {
			assert.Equal(t, maxDepth(tt.args.root), tt.want)
		})
	}
}
