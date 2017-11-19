package main

import "testing"

func TestReverseInt(t *testing.T) {
	if v := reverseInt(98); v != 89 {
		t.Errorf("Expected 89 but got %v", v)
	}
}
func TestHandleZero(t *testing.T) {
	if v := reverseInt(0); v != 0 {
		t.Errorf("Expected 0 but got %v", v)
	}
}

func TestFlipsAPositiveNumber(t *testing.T) {
	if v := reverseInt(5); v != 5 {
		t.Errorf("Expected 5 but got %v", v)
	}
	if v := reverseInt(15); v != 51 {
		t.Errorf("Expected 51 but got %v", v)
	}
	if v := reverseInt(89); v != 98 {
		t.Errorf("Expected 98 but got %v", v)
	}
	if v := reverseInt(2839); v != 9382 {
		t.Errorf("Expected 9382 but got %v", v)
	}
}

func TestHandleNegativeNumber(t *testing.T) {
	if v := reverseInt(-15); v != -51 {
		t.Errorf("Expected -51 but got %v", v)
	}
}

func TestHandleTrailingZeros(t *testing.T) {
	if v := reverseInt(1500); v != 51 {
		t.Errorf("Expected 51 but got %v", v)
	}
	if v := reverseInt(-1500); v != -51 {
		t.Errorf("Expected -51 but got %v", v)
	}
}
