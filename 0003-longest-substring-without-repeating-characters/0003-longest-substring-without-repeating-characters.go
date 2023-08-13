func lengthOfLongestSubstring(s string) int {
	longest := 0
	n := len(s)
	for i := 0; i < n && n-i > longest; i++ {
		m := make(map[byte]bool)
		for j := i; j < n; j++ {
			if m[s[j]] {
				break
			}

			m[s[j]] = true
		}

		currentLen := len(m)
		if currentLen > longest {
			longest = currentLen
		}
	}

	return longest
}