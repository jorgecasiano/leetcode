func isValid(s string) bool {
	if len(s) % 2 == 1 {
        return false
    }
	symbolsRune := []rune("[](){}")
	symbols := map[rune]rune{
		symbolsRune[1]: symbolsRune[0],
		symbolsRune[3]: symbolsRune[2],
		symbolsRune[5]: symbolsRune[4],
	}

	var stack []rune
	for _, symbol := range s {
		val, found := symbols[symbol]
		if found {
			if len(stack) == 0 || stack[last(stack)] != val {
				return false
			}

			stack = stack[:last(stack)]
		} else {
			stack = append(stack, symbol)
		}
	}

	return len(stack) == 0
}

func last(arr []rune) int {
	return len(arr) - 1
}