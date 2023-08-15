func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2
	if len(b) < len(a) {
		a, b = b, a
	}

	l, r := 0, len(a)
	half := (len(a) + len(b) + 1) / 2

	for l < r {
		i := l + (r-l)/2
		j := half - i

		if a[i] < b[j-1] {
			l = i + 1
		} else {
			r = i
		}
	}

	median1 := max(value(a, l-1, math.MinInt), value(b, half-l-1, math.MinInt))

	if (len(a)+len(b))%2 == 1 {
		return float64(median1)
	}

	median2 := min(value(a, l, math.MaxInt), value(b, half-l, math.MaxInt))
	return float64(median1+median2) * 0.5
}

func value(nums []int, index int, def int) int {
	if index >= 0 && index < len(nums) {
		return nums[index]
	}

	return def
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}