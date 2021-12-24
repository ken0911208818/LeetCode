package sortstring

import (
	"fmt"
	"sort"
)

func sortString(input ...string) []string {
	sort.Strings(input)
	for _, v := range input {
		fmt.Println([]byte(v[:1]))
		fmt.Println(v[:1])
		// if []byte(v[:1])[0] > byte(90) {

		// 	result := append(input[i:], input[:i]...)
		// 	return result
		// }
	}
	return input
}
