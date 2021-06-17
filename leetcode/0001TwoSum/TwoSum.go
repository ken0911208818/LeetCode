package twosum

func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, v := range nums {
        if k, ok := m[target - v]; ok {
			if i > k {
				return []int{k, i}
			}
            return []int{i, k}
        }
		m[v] = i 
    }
    return []int{}
}

