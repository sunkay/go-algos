package main

import "testing"

func TestPalindrome(t *testing.T) {
	if p := palindrome("abba"); p != true {
		t.Errorf("Expected abba == true but got %v", p)
	}

	if p := palindrome("aba "); p != false {
		t.Errorf("Expected aba  == false but got %v", p)
	}

	if p := palindrome("greetings"); p != false {
		t.Errorf("Expected greetings == false but got %v", p)
	}

	if p := palindrome("1000000001"); p != true {
		t.Errorf("Expected 1000000001 == true but got %v", p)
	}

	if p := palindrome("Fish hsif"); p != false {
		t.Errorf("Expected Fish hsif == false but got %v", p)
	}
}
func TestPalindromeArray(t *testing.T) {
	if p := palindromeArray("abba"); p != true {
		t.Errorf("Expected abba == true but got %v", p)
	}

	if p := palindromeArray("aba "); p != false {
		t.Errorf("Expected aba  == false but got %v", p)
	}

	if p := palindromeArray("greetings"); p != false {
		t.Errorf("Expected greetings == false but got %v", p)
	}

	if p := palindromeArray("1000000001"); p != true {
		t.Errorf("Expected 1000000001 == true but got %v", p)
	}

	if p := palindromeArray("Fish hsif"); p != false {
		t.Errorf("Expected Fish hsif == false but got %v", p)
	}

	if p := palindromeArray("pennep"); p != true {
		t.Errorf("Expected pennep == true but got %v", p)
	}
}
