func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs) - 1]
	min := int(math.Min(float64(len(first)), float64(len(last))))

	for i := 0; i < min; i++ {
		if first[i] != last[i] {
			return first[0:i]
		}
	}

	return strs[0]
}