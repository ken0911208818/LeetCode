package singleNumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run("singleNumber", func(t *testing.T) {
			fmt.Println("----------------136. Single Number---------------------")
			assert.Equal(t, singleNumber(tt.args.nums), tt.want)
			fmt.Printf("Input: %v\n", tt.args.nums)
			fmt.Printf("Output: %v \n", tt.want)
		})
	}
}
