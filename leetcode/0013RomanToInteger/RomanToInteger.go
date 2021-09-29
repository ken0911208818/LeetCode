package romanToInteger

func romanToInt(s string) int {
	mapValue := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0

	for i:= len(s) -1; i >= 0; i-- {
		key := string(s[i])
		lastkey := ""
		if i != 0 {
			lastkey = string(s[i-1])
		}
		value := mapValue[key]

		if mapValue[lastkey] < mapValue[key] {
			value = mapValue[key] - mapValue[lastkey]
			i--
		}
		result += value
	}
	return result
}
