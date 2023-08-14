func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
	len2 := len(nums2)
	totalSize := len1 + len2

	allNums := make([]int, totalSize)
	var i, j, k int
	for i <= len1 && j <= len2 {
		if i == len1 {
			for j < len2 {
				allNums[k] = nums2[j]
				k++
				j++
			}
			break
		} else if j == len2 {
			for i < len1 {
				allNums[k] = nums1[i]
				k++
				i++
			}
			break
		}

		if nums1[i] < nums2[j] {
			allNums[k] = nums1[i]
			i++
		} else {
			allNums[k] = nums2[j]
			j++
		}
		k++
	}

	if totalSize % 2 == 0 { return float64(allNums[totalSize/2-1] + allNums[totalSize/2]) / 2 }
	return float64(allNums[totalSize/2])
}