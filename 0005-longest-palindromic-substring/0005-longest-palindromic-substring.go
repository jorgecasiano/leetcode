func longestPalindrome(s string) string {
	sBytes := []byte(s)
	var res []byte

	for i := 0; i < len(sBytes); i++ {
		resEven := findPalindrome(sBytes, i, i)
		resOdd := findPalindrome(sBytes, i, i + 1)

		if len(resEven) > len(res) {
			res = resEven
		}

		if len(resOdd) > len(res) {
			res = resOdd
		}
	}

	return string(res)
}

func findPalindrome(s []byte, i int, right int) []byte {
	var res []byte
	l, r := i, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if (r - l + 1) > len(res) {
			res = s[l:r+1]
		}
		l--
		r++
	}

	return res
}