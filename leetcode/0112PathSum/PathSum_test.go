package pathSum

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	null := structures.NULL
	type args struct {
		root      []int
		targetSum int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				root:      []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1},
				targetSum: 22,
			},
			want: true,
		},
		{
			args: args{
				root:      []int{},
				targetSum: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("-------------------0112PathSum--------------------------", func(t *testing.T) {
			node := structures.Ints2TreeNode(tt.args.root)
			assert.Equal(t, hasPathSum(node, tt.args.targetSum), tt.want)
			fmt.Printf("Input: root = %v, targetSum = %v \n", tt.args.root, tt.args.targetSum)
			fmt.Printf("Output: %v", tt.want)
		})
	}
}
