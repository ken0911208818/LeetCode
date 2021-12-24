package sortstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortString(t *testing.T) {
	type args struct {
		in0 []string
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				in0: []string{"Apple", "apple", "zzzz", "ZZZZ"},
			},
			want: []string{"apple", "zzzz", "Apple", "ZZZZ"},
		},
		{
			args: args{
				in0: []string{"apple", "zzzz", "Apple", "ZZZZ"},
			},
			want: []string{"apple", "zzzz", "Apple", "ZZZZ"},
		},
		{
			args: args{
				in0: []string{"apple", "zzzz", "cccc", "dddd"},
			},
			want: []string{"apple", "cccc", "dddd", "zzzz"},
		},
		{
			args: args{
				in0: []string{"Apple", "Zzzz", "Cccc", "Dddd"},
			},
			want: []string{"Apple", "Cccc", "Dddd", "Zzzz"},
		},
	}
	for _, tt := range tests {
		t.Run("sort string", func(t *testing.T) {
			assert.Equal(t, sortString(tt.args.in0...), tt.want)
		})
	}
}
