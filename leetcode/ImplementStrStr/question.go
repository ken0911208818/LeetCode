package implementstr

import "strings"

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i > len(haystack)-len(needle) {
			return -1
		}
		if haystack[i] == needle[0] {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}

func strStr01(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr02(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
