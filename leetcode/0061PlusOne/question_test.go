package plusone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			args: args{
				digits: []int{0},
			},
			want: []int{1},
		},
		{
			args: args{
				digits: []int{9, 9},
			},
			want: []int{1, 0, 0},
		},
		{
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run("test Plus one", func(t *testing.T) {
			assert.Equal(t, plusOne(tt.args.digits), tt.want)
		})
	}
}
