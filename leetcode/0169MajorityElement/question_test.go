package majorityElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("majority number", func(t *testing.T) {
			assert.Equal(t, majorityElement(tt.args.nums), tt.want)
		})
	}
}
