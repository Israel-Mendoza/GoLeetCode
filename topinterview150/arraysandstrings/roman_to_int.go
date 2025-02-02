package arraysandstrings

var romanIntegerMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {

	lastValue := romanIntegerMap[s[len(s)-1]]
	sum := lastValue

	for i := len(s) - 2; i >= 0; i-- {
		currentValue := romanIntegerMap[s[i]]

		if currentValue < lastValue {
			currentValue *= -1
		}

		sum += currentValue
		lastValue = currentValue
	}

	return sum
}
