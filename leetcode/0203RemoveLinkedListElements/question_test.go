package removelinkedlistElement

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Input: head = [1,2,6,3,4,5,6], val = 6 Output: [1,2,3,4,5]",
			args: args{
				head: []int{1,2,6,3,4,5,6},
				val:  6,
			},
			want: []int{1,2,3,4,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, structures.List2Ints( removeElements(structures.Ints2List(tt.args.head), tt.args.val)), tt.want)
		})
	}
}
