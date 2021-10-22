package pascalTriangle

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		gen := make([]int, i)
		gen[0] = 1
		if len(gen) >= 2 {
			gen[i-1] = 1
		}
		for j := 1; j < i-1; j++ {
			if i == 1 {
				continue
			}
			gen[j] = result[i-2][j-1] + result[i-2][j]
		}
		result[i-1] = gen
	}
	return result
}

func generate01(numRows int) [][]int {
	result := [][]int{}
	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result
}