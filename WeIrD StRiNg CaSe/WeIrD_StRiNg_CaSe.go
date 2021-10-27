package main

import "strings"

func toWeirdCase(text string) string {
	var wiredCaseText string
	for _, word := range strings.Split(text, " "){
		if wiredCaseText != ""{
			wiredCaseText += " "
		}
		var wiredCaseWord string
		for index, digitCode := range word{
			digit := string(digitCode)
			if index % 2 == 0{
				digit = strings.ToUpper(digit)
			} else {
				digit = strings.ToLower(digit)
			}
			wiredCaseWord += digit
		}
		wiredCaseText += wiredCaseWord
	}

	return wiredCaseText
}