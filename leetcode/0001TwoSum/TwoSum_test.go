package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	num    []int
	target int
}

type ans1 struct {
	ans1 []int
}

func TestTwoSum(t *testing.T) {
	qs := []question1{
		{
			para1: para1{[]int{3, 2, 4}, 6},
			ans1:  ans1{[]int{1, 2}},
		},
		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
	}
	for _, q := range qs {
		a, p := q.ans1, q.para1

	assert.Equal(t, twoSum(p.num, p.target), a.ans1)
	}
}
