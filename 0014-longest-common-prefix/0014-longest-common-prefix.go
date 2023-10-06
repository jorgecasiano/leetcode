func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	pos := 0
	for pos < len(strs[0]) {
		for i := 1; i < len(strs); i++ {
			if pos == len(strs[i-1]) || pos == len(strs[i]) || strs[i-1][pos] != strs[i][pos] {
				return strs[0][0:pos]
			}

			if i == len(strs)-1 {
				pos++
			}
		}
	}

	return strs[0]
}