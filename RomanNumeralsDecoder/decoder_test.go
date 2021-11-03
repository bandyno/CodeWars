package RomanNumeralsDecoder

import (
	"testing"
)

func TestDecodeOneDigit(t *testing.T){
	// Arrange
	romanNumber := "I"
	decimalNumber := 1

	// Act
	result := Decode(romanNumber)

	// Assert
	if result != decimalNumber {
		t.Errorf("func result was incorrect, got: %v, want: %v", result, decimalNumber)
	}
}
func TestDecodeFewDigits(t *testing.T)  {
	// arrange
	romanNumber := "XVII"
	decimalNumber := 17

	romanNumber1 := "XIV"
	decimalNumber1 := 14
	// act
	result := Decode(romanNumber)
	result1 := Decode(romanNumber1)
	// assert
	if decimalNumber != result{
		t.Errorf("func result was incorrect, got: %v, want: %v", result, decimalNumber)
	} 

	if decimalNumber1 != result1{
		t.Errorf("func result was incorrect, got: %v, want: %v", result1, decimalNumber1)
	}

}
