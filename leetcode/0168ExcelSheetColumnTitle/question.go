package excelsheetcolumntitle



func convertToTitle(columnNumber int) string {
    result :=[]byte{}
	for columnNumber > 0 {
		result = append(result, 'A'+byte((columnNumber-1)%26))
		columnNumber = (columnNumber-1)/26
	}
	return reverseStrRecursion(string(result))
}

func reverseStrRecursion(input string) string {
	runes := []byte(input)
	if len(runes) <= 1 {
	  return string(runes)
	}
	return reverseStrRecursion(string(runes[1:])) + string(runes[0])
}