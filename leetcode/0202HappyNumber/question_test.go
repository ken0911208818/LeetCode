package happynumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input: 19 output: 82 ",
			args: args{
				n: 19,
			},
			want: 82,
		},
		{
			name: "input: 82 output: 68 ",
			args: args{
				n: 82,
			},
			want: 68,
		},
		{
			name: "input: 100 output: 1 ",
			args: args{
				n: 100,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, calculate(tt.args.n), tt.want)
		})
	}
}

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "input 19 is happyNumber",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "input 2 is happyNumber",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "input 7 is happyNumber",
			args: args{
				n: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, isHappy(tt.args.n),tt.want)
		})
	}
}
