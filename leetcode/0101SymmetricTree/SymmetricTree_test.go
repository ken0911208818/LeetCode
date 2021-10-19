package symmetricTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2	,
						Left:  &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2	,
						Left:  &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("SymmetricTree", func(t *testing.T) {
			assert.Equal(t, isSymmetric(tt.args.root), tt.want)
		})
	}
}
