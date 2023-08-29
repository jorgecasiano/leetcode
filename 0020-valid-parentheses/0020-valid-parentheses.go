const symbols = "[({])}"

func isValid(s string) bool {
	openSymbols := map[byte]int{
		symbols[0]: 0,
		symbols[1]: 1,
		symbols[2]: 2,
	}

	closeSymbols := map[byte]int{
		symbols[3]: 0,
		symbols[4]: 1,
		symbols[5]: 2,
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		val, found := closeSymbols[s[i]]
		if found {
			if len(stack) == 0 || openSymbols[stack[len(stack)-1]] != val {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			_, found = openSymbols[s[i]]
			if found {
				stack = append(stack, s[i])
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}