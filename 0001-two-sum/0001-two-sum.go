func twoSum(nums []int, target int) []int {
	processed := make(map[int]int)
	for i, num := range nums {
		if j, found := processed[target-num]; found {
			return []int{i,j}
		}
		processed[num] = i
	}
	panic("Solution not found")
}