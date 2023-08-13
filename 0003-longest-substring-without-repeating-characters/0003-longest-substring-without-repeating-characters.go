func lengthOfLongestSubstring(s string) int {
	lastIndexChar := make(map[rune]int)
	longestSub := 0
	currentSub := 0
	start := 0
	
	for index, char := range s {
		lastIndex, found := lastIndexChar[char]
		if !found || lastIndex < index - currentSub {
			currentSub++
		} else {
			if currentSub > longestSub {
				longestSub = currentSub
			}

			start = lastIndex + 1
			currentSub = index - start + 1
		}
		lastIndexChar[char] = index
	}

	if currentSub > longestSub {
		longestSub = currentSub
	}

	return longestSub
}