package addbinary

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	i, j := len(a), len(b)
	c := 0
	res := make([]string, len(a)+1)
	for i > 0 && j > 0 {
		ai, _ := strconv.Atoi(string(a[i-1]))
		bj, _ := strconv.Atoi(string(b[j-1]))
		res[i] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
	}

	for i > 0 {
		ai, _ := strconv.Atoi(string(a[i-1]))
		res[i] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) /2
		i --
	}
	if c > 0 {
		res[0] = strconv.Itoa(c)
	}
	return strings.Join(res, "")
}
