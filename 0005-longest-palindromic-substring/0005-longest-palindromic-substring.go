func longestPalindrome(s string) string {
	var l, r, maxLen int

	for i := 0; i < len(s); i++ {
		lEven, rEven := findPalindrome(s, i, i)
		lOdd, rOdd := findPalindrome(s, i, i + 1)

		currentLength := calcLength(lEven, rEven)
		if currentLength > maxLen {
			maxLen = currentLength
			l, r = lEven, rEven
		}

		currentLength = calcLength(lOdd, rOdd)
		if currentLength > maxLen {
			maxLen = currentLength
			l, r = lOdd, rOdd
		}
	}

	return string(s[l:r+1])
}

func findPalindrome(s string, i int, right int) (int, int) {
	l, r := i, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return l + 1, r - 1
}

func calcLength(l int, r int) int {
	return r - l + 1
}