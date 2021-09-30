package sum3

import "sort"

func threeSum(nums []int) [][]int {
	// -4, -1, -1, 0, 1, 2
	sort.Ints(nums)
	start, end, result := 0, len(nums)-1, make([][]int, 0)
	for i := 1; i < len(nums)-1; i++ {
		start, end = 0, len(nums)-1
		if i > 1 && nums[i] == nums[i-1] {
			start = i - 1
		}

		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum := nums[start] + nums[end] + nums[i]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			}
			if addNum < 0 {
				start++
			}

			if addNum > 0 {
				end--
			}
		}
	}
	
	return result
}
