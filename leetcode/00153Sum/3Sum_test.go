package sum3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	ans
	para
}

type ans struct {
	ans [][]int
}

type para struct {
	nums []int
}

func Test_threeSum(t *testing.T) {
	qs := []question{
		{
			ans: ans{
				ans: [][]int{{-1, -1, 2}, {-1, 0, 1}},
			},
			para: para{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
		},
	}

	for _, v := range qs {
		p, a := v.nums, v.ans
		assert.Equal(t, threeSum(p), a.ans)
	}
}
