package RomanNumeralsDecoder

func Decode(roman string) int {
	dict :=  make(map[string]int)
	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10

	result := 0
	for _, v := range roman{
		result += dict[string(v)]
	}

	return result
}