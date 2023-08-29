func romanToInt(s string) int {
	symbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	romanNumber := []rune(s)

	if len(romanNumber) == 1 {
		return symbols[romanNumber[0]]
	}

	var i, acc int

	for i < len(romanNumber) - 1 {
		current := symbols[romanNumber[i]]
		next := symbols[romanNumber[i + 1]]

		if current >= next {
			acc += current
			i++
		} else {
			acc += next - current
			i += 2
		}
	}

	if i == len(romanNumber) {
		return acc
	}

	last := symbols[romanNumber[len(romanNumber) - 1]]
	return acc + last
}