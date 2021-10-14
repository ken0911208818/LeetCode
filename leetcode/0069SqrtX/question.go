package sqrtX

func mySqrt(x int) int {
	res := 1

	for res*res < x {
		res += 1
	}
	if res*res == x {
		return res
	}
	return res - 1
}

// 使用牛頓求根法 https://www.youtube.com/watch?v=Quw4ZHLH2CY
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
