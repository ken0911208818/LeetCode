package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	input
	output
}

type input struct {
	s string
}

type output struct {
	ans bool
}

func Test_isValid(t *testing.T) {
	qs := []question{
		{
			input{"()"},
			output{true},
		},
		{
			input{"()[]{}"},
			output{true},
		},
		{
			input{"(]"},
			output{false},
		},
		{
			input{"([)]"},
			output{false},
		},
		{
			input{"{[]}"},
			output{true},
		},
	}

	for _, v := range qs {
		assert.Equal(t, isValid(v.s), v.ans)
	}
}

func Benchmark01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		isValid01("{}")
	}
}

func BenchmarkSelf(t *testing.B) {
	for i := 0; i < t.N; i++ {
		isValid("{}")
	}
}
