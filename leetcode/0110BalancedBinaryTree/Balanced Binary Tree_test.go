package balancedBinaryTree

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	null := structures.NULL
	type args struct {
		root []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				root: []int{3, 9, 20, null, null, 15, 7},
			},
			want: true,
		},
		{
			args: args{
				root: []int{1, 2, 2, 3, 3, null, null, 4, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		node := structures.Ints2TreeNode(tt.args.root)
		assert.Equal(t, isBalanced(node), tt.want)
		fmt.Printf("input: %v	", tt.args.root)
		fmt.Printf("output: %v	\n", tt.want)
	}
}
