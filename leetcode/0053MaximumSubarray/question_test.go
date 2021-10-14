package maximumsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{nums: []int{-2,1,-3,4,-1,2,1,-5,4}},
			want: 6,
		},
		{
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			args: args{nums: []int{5,4,-1,7,8}},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run("maxSubArray", func(t *testing.T) {
			assert.Equal(t, maxSubArray(tt.args.nums), tt.want)
		})
	}
}
