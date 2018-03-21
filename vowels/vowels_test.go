package vowels

import (
	"testing"
)

func TestVowels(t *testing.T) {

	if v := vowels("aEIOU"); v != 5 {
		t.Errorf("Expected number to be %v, but got %v", 5, v)
	}

	if v := vowels("abcdefghijklmnopqrstuvwxyz"); v != 5 {
		t.Errorf("Expected number to be %v, but got %v", 5, v)
	}

	if v := vowels("bcdfghjkl"); v != 0 {
		t.Errorf("Expected number to be %v, but got %v", 0, v)
	}
}
