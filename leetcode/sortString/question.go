package sortstring

import (
	"sort"
)

func sortString(input ...string) []string {
	sort.Strings(input)
	for i, v := range input {
		if []byte(v[:1])[0] > byte(90) {

			result := append(input[i:], input[:i]...)
			return result
		}
	}
	return input
}
