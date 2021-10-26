package main

import (
	"testing"
)

func TestToWeirdCase(t *testing.T) {
	// arrange
	text := "SoMe StRiNg"

	// act
	result := toWeirdCase("some string")

	// assert
	if result != text {
		t.Errorf("Case was incorrected, got: %v, want: %v", text, result)

	}
}
