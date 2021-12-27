package excelsheetcolumnnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Title : A output :1",
			args: args{
				columnTitle: "A",
			},
			want: 1,
		},
		{
			name: "Title : AB output :28",
			args: args{
				columnTitle: "AB",
			},
			want: 28,
		},
		{
			name: "Title : ZY output :701",
			args: args{
				columnTitle: "ZY",
			},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, titleToNumber(tt.args.columnTitle), tt.want)
		})
	}
}
