func removeElement(nums []int, val int) int {
    var k int
    r := len(nums) - 1
    
    for k <= r {
        if val != nums[k] {
            k++
        } else {
            nums[k], nums[r] = nums[r], nums[k]
            r--
        }
    }
    
    return k 
}