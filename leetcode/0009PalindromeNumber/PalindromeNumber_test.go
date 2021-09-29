package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	num int
}

type ans1 struct {
	ans1 bool
}

func Test_isPalindrome(t *testing.T) {
	qs := []question1{
		{
			para1: para1{121},
			ans1:  ans1{true},
		},
		{
			para1: para1{-121},
			ans1:  ans1{false},
		},
		{
			para1: para1{10},
			ans1:  ans1{false},
		},
		{
			para1: para1{-101},
			ans1:  ans1{false},
		},
	}

	for _, q := range qs {
		a, p := q.ans1, q.para1
		assert.Equal(t, isPalindrome(p.num), a.ans1)
	}
}
