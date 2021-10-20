package convertSortedArraytoBinarySearchTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				nums: []int{0,-3,9,-10,null,5},
			},
		},
	}
	for _, tt := range tests {
		t.Run("convertSortedArraytoBinarySearchTree", func(t *testing.T) {
			var result []int
			want := Preorder(sortedArrayToBST(tt.args.nums), &result)
			assert.Equal(t, want, tt.args)
		})
	}
}
