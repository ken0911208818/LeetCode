package majorityElement

func majorityElement(nums []int) int {
    result := make(map[int]int)
	for _, v := range nums {
		result[v] +=1
		if len(nums)/2 < result[v] {
			return v
		}
	}
	return 0
}