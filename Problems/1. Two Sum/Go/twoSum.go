func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i, n := range nums {
        if idx, ok := numsMap[target - n]; ok {
            return []int{idx, i}
        }
        numsMap[n] = i
    }
    return []int{}
}
