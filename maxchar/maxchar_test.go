package main

import "testing"

func TestMostFreqChar(t *testing.T) {
	if f := maxChar("a"); f != "a" {
		t.Errorf("Expected a got %v", f)
	}

	if f := maxChar("abcdefghijklmnaaaaa"); f != "a" {
		t.Errorf("Expected a got %v", f)
	}
}

func TestMostFreqCharsWithNumbers(t *testing.T) {
	if f := maxChar("ab1c1d1e1f1g1"); f != "1" {
		t.Errorf("Expected 1 got %v", f)
	}
}
