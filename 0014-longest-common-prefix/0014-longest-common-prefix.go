func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	pos := 0
	out:
	for pos < len(strs[0]) {
		for i := 1; i < len(strs); i++ {
			if pos == len(strs[i-1]) || pos == len(strs[i]) {
				break out
			}
			if strs[i-1][pos] != strs[i][pos] {
				break out
			}

			if i == len(strs)-1 {
				pos++
			}
		}
	}

	return strs[0][0:pos]
}