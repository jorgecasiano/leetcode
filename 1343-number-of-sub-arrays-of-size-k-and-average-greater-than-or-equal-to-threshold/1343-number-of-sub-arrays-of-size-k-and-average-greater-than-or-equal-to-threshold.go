func numOfSubarrays(arr []int, k int, threshold int) int {
    var i, j, acc, sum int

	for j < len(arr) {
		sum += arr[j]
		if j - i + 1 == k {
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