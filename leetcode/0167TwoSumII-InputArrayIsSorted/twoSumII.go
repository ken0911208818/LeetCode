package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
    mapNumbers := make(map[int]int)  
    for i:=0; i<len(numbers); i++ {
        number := numbers[i]
        if v, has := mapNumbers[target-number]; has {
            return []int{v, i+1}
        }
        mapNumbers[number] = i+1
        
    }
    return []int{0,0}
}