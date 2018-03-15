package vowels

import (
	"testing"
)

func TestVowels(t *testing.T) {

	if v := vowels("AEIOU"); v != 5 {
		t.Errorf("Expected number to be %v, but got %v", 5, v)
	}
}
