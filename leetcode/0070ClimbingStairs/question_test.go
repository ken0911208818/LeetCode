package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{n: 2},
			want: 2,
		},
		{
			args: args{n: 3},
			want: 3,
		},
		{
			args: args{n: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run("climbingstairs", func(t *testing.T) {
			assert.Equal(t, climbStairs(tt.args.n), tt.want)
		})
	}
}
