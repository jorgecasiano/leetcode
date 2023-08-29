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

	size, maxSize := 0, len(s) / 2
	stack := make([]rune, maxSize)
	for _, symbol := range s {
		if size < 0 || size > maxSize  {
			return false
		}
		val, found := symbols[symbol]
		if found {
			if size == 0 || stack[size - 1] != val {
				return false
			}

			size--
		} else if size < maxSize {            
			stack[size] = symbol
			size++
		} else {
			return false
		}
	}

	return size == 0
}