package anagrams

import "testing"

func TestAnaHello(t *testing.T) {
	testAnagram("Hello", "olleH", true, t)
}

func TestAnaWoah(t *testing.T) {
	testAnagram("Whoa! Hi", "Hi! Whoa!", true, t)
}

func TestAnaOneOne(t *testing.T) {
	testAnagram("One One", "Two two two", false, t)
}

func TestAnaOneOnec(t *testing.T) {
	testAnagram("One one", "One one c", false, t)
}

func TestAnaPunctuation(t *testing.T) {
	testAnagram("A tree, a life, a bench", "A tree, a fence, a yard", false, t)
}

func testAnagram(strA string, strB string, expected bool, t *testing.T) {
	r := anagramSortSolution(strA, strB)
	if r != expected {
		t.Errorf("strA: %v, strB: %v, Expected %v, but got %v", strA, strB, expected, r)
	}

	r1 := anagram(strA, strB)
	if r1 != expected {
		t.Errorf("strA: %v, strB: %v, Expected %v, but got %v", strA, strB, expected, r1)
	}
}
