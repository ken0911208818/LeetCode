package implementstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	input
	output
}

type input struct {
	haystack string
	needle   string
}

type output struct {
	result int
}

var BenchmarkValue = input{
	haystack: "hello",
	needle:   "ll",
}

func Test_strStr(t *testing.T) {
	qs := []question{
		{
			input{
				haystack: "hello",
				needle:   "ll",
			},
			output{
				result: 2,
			},
		},
		{
			input{
				haystack: "",
				needle:   "",
			},
			output{
				result: 0,
			},
		},
		{
			input{
				haystack: "aaaaaaa",
				needle:   "bba",
			},
			output{
				result: -1,
			},
		},
	}
	for _, v := range qs {
		assert.Equal(t, strStr(v.haystack, v.needle), v.result)
	}
}


func BenchmarkSelf(t *testing.B) {
	for i:= 0; i< t.N; i++ {
		strStr(BenchmarkValue.haystack, BenchmarkValue.needle)
	}
}

func Benchmark01(t *testing.B) {
	for i:= 0; i< t.N; i++ {
		strStr01(BenchmarkValue.haystack, BenchmarkValue.needle)
	}
}

func Benchmark02(t *testing.B) {
	for i:= 0; i< t.N; i++ {
		strStr02(BenchmarkValue.haystack, BenchmarkValue.needle)
	}
}