import "math"

func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}

	isNegative := x < 0

	if isNegative {
		x *= -1
	}

    var val int

	for x > 9 {
		val += x % 10
		val *= 10

        if val > math.MaxInt32 {
            return 0
        }

		x /= 10
	}

    val += x

    if val > math.MaxInt32 {
        return 0
    }
    
	if isNegative {
		val *= -1
	}

    return val
	
}