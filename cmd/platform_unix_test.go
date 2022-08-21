package cmd

import (
	"testing"
)

func TestUnixNoSpecial(t *testing.T) {
	input := "no special characters"
	expected := "'no special characters'"
	result := sanitizeMessage(input)
	if result != expected {
		t.Errorf("Expected %x; got %x", expected, result)
	}
}
