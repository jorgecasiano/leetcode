func numOfSubarrays(arr []int, k int, threshold int) int {
	j := k - 1
	sum := sumFirstN(arr, j)

	var i, acc int

	for j < len(arr) {
		sum += arr[j]

		if sum / k >= threshold {
			acc++
		}

		sum -= arr[i]
		i++
		j++
	}

	return acc
}

func sumFirstN(arr []int, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	return sum
}