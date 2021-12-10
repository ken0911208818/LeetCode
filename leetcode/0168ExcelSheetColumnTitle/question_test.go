package excelsheetcolumntitle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				columnNumber: 1,
			},
			want: "A",
		},
		{
			args: args{
				columnNumber: 28,
			},
			want: "AB",
		},
		{
			args: args{
				columnNumber: 701,
			},
			want: "ZY",
		},
		{
			args: args{
				columnNumber: 2147483647,
			},
			want: "FXSHRXW",
		},
	}
	for _, tt := range tests {
		t.Run("convert number to string", func(t *testing.T) {
			assert.Equal(t, convertToTitle(tt.args.columnNumber), tt.want)
		})
	}
}
