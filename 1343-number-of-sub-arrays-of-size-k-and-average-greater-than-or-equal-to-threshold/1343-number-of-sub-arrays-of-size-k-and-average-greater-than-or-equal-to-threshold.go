func numOfSubarrays(arr []int, k int, threshold int) int {
    var i, j, acc, sum int
	var fullWindow = false
    
	for j < len(arr) {
		sum += arr[j]
		if fullWindow || (j - i + 1 == k) {
			fullWindow = true
			if sum / k >= threshold {
				acc++
			}

			sum -= arr[i]
			i++
		}

		j++
	}

	return acc
}