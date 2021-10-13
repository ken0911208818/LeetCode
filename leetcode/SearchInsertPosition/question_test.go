package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	input
	output
}

type input struct {
	nums   []int
	target int
}

type output struct {
	result int
}

func Test_searchInsert(t *testing.T) {
	qs := []question{
		{
			input{nums: []int{1, 3, 5, 6}, target: 5},
			output{result: 2},
		},
		{
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			output: output{
				result: 1,
			},
		},
		{
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			output: output{
				result: 4,
			},
		},
		{
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			output: output{
				result: 0,
			},
		},
	}

	for _, v := range qs {
		assert.Equal(t, searchInsert(v.nums, v.target), v.result)
	}
}
