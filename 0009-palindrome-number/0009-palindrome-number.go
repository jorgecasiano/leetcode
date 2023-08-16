func isPalindrome(x int) bool {
	// exclude negatives, zero and numbers ending in zero
	if x < 0 || (x != 0 && x % 10 == 0) {
		return false
	}

	reverse := 0
	for x > reverse {
		reverse = reverse * 10 + x % 10;
		x = x / 10;
	}

	// x == reverse -> for even palindromes
	// x == reverse / 10 -> for odd palindromes
	return x == reverse || x == reverse / 10;
}