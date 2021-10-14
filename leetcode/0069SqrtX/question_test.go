package sqrtX

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{x: 101},
			want: 10,
		},
		{
			args: args{x: 4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run("SqrtX", func(t *testing.T) {
			assert.Equal(t, mySqrt1(tt.args.x), tt.want)
		})
	}
}

func BenchmarkSelf(t *testing.B) {
	for i:= 0; i< t.N; i++ {
		mySqrt(2555555)
	}
}

func Benchmark01(t *testing.B) {
	for i:= 0; i < t.N; i++ {
		mySqrt1(2555555)
	}
}
