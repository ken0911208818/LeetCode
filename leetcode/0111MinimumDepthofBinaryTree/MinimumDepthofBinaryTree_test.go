package minimumDepthOfBinaryTree

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_minDepth(t *testing.T) {
	null := structures.NULL
	type args struct {
		root []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				root: []int{3, 9, 20, null, null, 15, 7},
			},
			want: 2,
		},
		{
			args: args{
				root: []int{2, null, 3, null, 4, null, 5, null, 6},
			},
			want: 5,
		},
		{
			args: args{
				root: []int{-9,-3,2,null,4,4,0,-6,null,-5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("minimumDepthOfBinaryTree", func(t *testing.T) {
			node := structures.Ints2TreeNode(tt.args.root)
			assert.Equal(t, minDepth(node), tt.want)
		})
		fmt.Println("-----------------111. Minimum Depth of Binary Tree--------------------------")
		fmt.Printf("input: %v	\n", tt.args.root)
		fmt.Printf("ouput: %v	\n", tt.want)
	}
}
