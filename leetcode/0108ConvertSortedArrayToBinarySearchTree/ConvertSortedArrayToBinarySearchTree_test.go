package convertSortedArraytoBinarySearchTree

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want args
	}{
		{
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			want: args{
				nums: []int{0, -3, -10, 9, structures.NULL, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run("convertSortedArraytoBinarySearchTree", func(t *testing.T) {
			var result []int
			structures.T2s(sortedArrayToBST(tt.args.nums), &result)
			fmt.Printf("input: %v        output: %v", tt.args, result)
		})
	}
}
