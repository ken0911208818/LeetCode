package removeElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	input
	output
}

type input struct {
	nums []int
	val  int
}

type output struct {
	ans int
}

func Test_removeElement(t *testing.T) {
	qs := []question{
		{
			input{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			output{
				ans: 2,
			},
		},
	}

	for _, v := range qs {
		assert.Equal(t, removeElement(v.nums, v.val), v.ans)
	}
}
