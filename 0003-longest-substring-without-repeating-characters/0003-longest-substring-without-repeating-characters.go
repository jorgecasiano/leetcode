func lengthOfLongestSubstring(s string) int {
	visited := make(map[byte]bool)
	left := 0
	longest := 0
	current := 0
	
	for right := 0; right < len(s); right++ {
		if found := visited[s[right]]; found {
			for {
				delete(visited, s[left])
				leftUpdated := s[left] == s[right]
				left++
				if leftUpdated {
					break
				}
			}
			longest = max(longest, current)
			current = right - left
		}

		visited[s[right]] = true
		current++
	}

	longest = max(longest, current)

	return longest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}