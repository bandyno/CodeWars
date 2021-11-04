package RomanNumeralsDecoder

func Decode(roman string) int {
	dict := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	result := 0
	lastNum := 0
	for _, v := range roman {
		number := dict[string(v)]
		if result == 0 {
			result += number
			lastNum = number
		} else {
			if number > lastNum {
				result += number - lastNum*2
			} else {
				result += number
			}
		}
		lastNum = number
	}

	return result
}
