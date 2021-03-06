package twosumiiinputarrayissorted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run("two Sum II", func(t *testing.T) {
			fmt.Printf("Input: numbers= %v, target = %v\n", tt.args.numbers, tt.args.target)
			fmt.Printf("output: %v\n", tt.want)
			assert.Equal(t, twoSum(tt.args.numbers, tt.args.target), tt.want)
		})
	}
}
