package pascalstriangleii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
		{
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
		{
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run("PascalsTriangleII", func(t *testing.T) {
			fmt.Println("------------119. Pascal's Triangle II-------------------")
			assert.Equal(t, getRow(tt.args.rowIndex), tt.want)
		})
	}
}
