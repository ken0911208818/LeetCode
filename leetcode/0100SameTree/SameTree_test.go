package sametree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				q: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				q: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
		{
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				q: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
		{
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				q: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("SameTree", func(t *testing.T) {
			assert.Equal(t, isSameTree(tt.args.p, tt.args.q), tt.want)
		})
	}
}
