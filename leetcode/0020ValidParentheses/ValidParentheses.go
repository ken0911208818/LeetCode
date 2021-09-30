package validparentheses

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
			continue
		}
		if len(stack) <= 0 {
			return false
		}

		if (v == '}' && stack[len(stack)-1] == '{') ||
			(v == ')' && stack[len(stack)-1] == '(') ||
			(v == ']' && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}


// example 
func isValid01(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
