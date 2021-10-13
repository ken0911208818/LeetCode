package searchinsertposition

func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    for i, v := range nums {
        if v >= target {
            return i
        }
    }
    return len(nums)
}