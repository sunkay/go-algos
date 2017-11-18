package main

import "testing"

func TestReverse(t *testing.T) {
	if reverse("abcd") != "dcba" {
		t.Errorf("Expected dcba but got %v", reverse("abcd"))
	}

	if reverse("  abcd") != "dcba  " {
		t.Errorf("Expected dcba__ but got %v", reverse("abcd  "))
	}
}
