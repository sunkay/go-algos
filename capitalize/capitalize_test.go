package capitalize

import "testing"

func TestCapitalize(t *testing.T) {
	testCapitalize("hey you", "Hey You", t)
	testCapitalize("i love you", "I Love You", t)
	testCapitalize("hi there, how is it going?", "Hi There, How Is It Going?", t)
}

func testCapitalize(orig string, expected string, t *testing.T) {
	r := capitalize(orig)
	if r != expected {
		t.Errorf("Original: %v, Expected %v, but got %v", orig, expected, r)
	}
}
