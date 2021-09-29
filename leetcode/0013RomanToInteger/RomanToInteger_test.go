package romanToInteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	num string
}

type ans1 struct {
	ans1 int
}

func Test_romanToInt(t *testing.T) {
	qs := []question1{
		{
			para1: para1{"III"},
			ans1:  ans1{3},
		},
		{
			para1: para1{"IV"},
			ans1:  ans1{4},
		},
		{
			para1: para1{"IX"},
			ans1:  ans1{9},
		},
		{
			para1: para1{"LVIII"},
			ans1:  ans1{58},
		},
		{
			para1: para1{"MCMXCIV"},
			ans1:  ans1{1994},
		},
	}
	for _, v := range qs {
		a, q := v.ans1, v.para1
		assert.Equal(t, romanToInt(q.num), a.ans1)
	}
}
