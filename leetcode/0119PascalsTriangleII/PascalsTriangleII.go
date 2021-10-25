package pascalstriangleii

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}
	result[0] = 1
	result[1] = 1

	for i := 2; i <= rowIndex; i++ {
		last := result[:i]
		result[i] = 1
		for j := i-1; j >= 1; j-- {
			result[j] = last[j-1] + last[j]
		}
	}
	return result
}
