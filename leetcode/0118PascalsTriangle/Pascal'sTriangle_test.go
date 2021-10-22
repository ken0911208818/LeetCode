package pascalTriangle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{
				numRows: 1,
			},
			want: [][]int{{1}},
		},
		{
			args: args{
				numRows: 5,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run("pascalTriangle", func(t *testing.T) {
			fmt.Println("-------------118. Pascal's Triangle-------------")
			assert.Equal(t, generate(tt.args.numRows), tt.want)
			fmt.Printf("Input: numRows = %v \n", tt.args.numRows)
			fmt.Printf("Output: %v \n", tt.want)
		})
	}
}

func BenchmarkSelf(t *testing.B) {
	for i:=0; i< t.N; i++ {
		generate(100)
	}
}

func Benchmark01(t *testing.B) {
	for i:=0; i< t.N; i++ {
		generate01(100)
	}
}
