package happynumber

func isHappy(n int) bool {
	record := map[int] struct{}{}
	for n != 1 {
		record[n] = struct{}{}
		n = calculate(n)
		if _, has := record[n]; has {
			return false
		}
	}
	return true
}

func calculate(n int) int {
	output := 0
	for n != 0 {
		rem := n % 10
		output += rem * rem
		n = n / 10
	}
	return output
}
