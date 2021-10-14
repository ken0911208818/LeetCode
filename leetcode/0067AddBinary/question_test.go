package addbinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{a: "11", b: "1"},
			want: "100",
		},
		{
			args: args{a: "1010", b: "1011"},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run("addBinary", func(t *testing.T) {
			assert.Equal(t, addBinary(tt.args.a, tt.args.b), tt.want)
		})
	}
}
