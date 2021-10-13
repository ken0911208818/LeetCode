package plusone

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			// no carry
			return digits
		}
		// carry
		digits[i] = 0
	}
	result := append([]int{1}, digits...)
	return result
}
