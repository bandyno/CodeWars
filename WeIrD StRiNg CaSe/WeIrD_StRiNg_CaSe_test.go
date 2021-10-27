package main

import (
	"testing"
)

func TestToWeirdCaseOneWordUpperCase(t *testing.T) {
	// arrange
	text := "SOME"
	wiredCaseText := "SoMe"

	// act
	result := toWeirdCase(text)

	// assert
	if result != wiredCaseText {
		t.Errorf("Case was incorrected, got: %v, want: %v", text, wiredCaseText)

	}
}
func TestToWeirdCaseOneWordLowerCase(t *testing.T) {
	// arrange
	text := "some"
	wiredCaseText := "SoMe"

	// act
	result := toWeirdCase(text)

	// assert
	if result != wiredCaseText {
		t.Errorf("Case was incorrected, got: %v, want: %v", text, wiredCaseText)

	}
}
func TestToWeirdCaseTwoWordsLowerCase(t *testing.T) {
	// arrange
	text := "some string"
	wiredCaseText := "SoMe StRiNg"

	// act
	result := toWeirdCase(text)

	// assert
	if result != wiredCaseText {
		t.Errorf("Case was incorrected, got: %v, want: %v", result, wiredCaseText)

	}
}
func TestToWeirdCaseTwoWordsUpperCase(t *testing.T) {
	// arrange
	text := "SOME STRING"
	wiredCaseText := "SoMe StRiNg"

	// act
	result := toWeirdCase(text)

	// assert
	if result != wiredCaseText {
		t.Errorf("Case was incorrected, got: %v, want: %v", result, wiredCaseText)

	}
}
