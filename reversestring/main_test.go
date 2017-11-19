package main

import "testing"

func TestReverse(t *testing.T) {
	if r := reverse("abcd"); r != "dcba" {
		t.Errorf("Expected dcba but got %v", r)
	}

	if r := reverse("  abcd"); r != "dcba  " {
		t.Errorf("Expected dcba__ but got %v", r)
	}
}
