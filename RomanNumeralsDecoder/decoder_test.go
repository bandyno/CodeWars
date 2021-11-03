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
	// act
	result := Decode(romanNumber)
	// assert
	if decimalNumber != result{
		t.Errorf("func result was incorrect, got: %v, want: %v", result, decimalNumber)
	}

}
