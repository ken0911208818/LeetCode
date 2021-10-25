package bestTimeToBuyandSellStock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			fmt.Println("-----------------121. Best Time to Buy and Sell Stock-------------------")
			fmt.Printf("Input: %v\n", tt.args.prices)
			fmt.Printf("Output: %v\n", tt.want)
			assert.Equal(t, maxProfit(tt.args.prices), tt.want)
		})
	}
}
