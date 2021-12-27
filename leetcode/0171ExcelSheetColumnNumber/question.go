package excelsheetcolumnnumber

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	// result := 0
	column := []byte(columnTitle)
	result := 0
	a := float64(0)
	for i := len(column) - 1; i >= 0; i-- {
		result += (int(column[i]) - 64) * int(math.Pow(26, a))
		a++
	}
	return result
}
