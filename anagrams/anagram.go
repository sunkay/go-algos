package anagrams

import (
	"regexp"
	"sort"
	"strings"
)

// --- Directions
// Check to see if two provided strings are anagrams of eachother.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//   anagrams('rail safety', 'fairy tales') --> True
//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//   anagrams('Hi there', 'Bye there') --> False

func anagram(strA string, strB string) bool {

	//`Compile` an optimized `Regexp` struct alphanumeric
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	charMapA := charMap(reg.ReplaceAllString(strA, ""))
	charMapB := charMap(reg.ReplaceAllString(strB, ""))

	// check lengths of 2 charMaps
	if len(charMapA) != len(charMapB) {
		return false
	}

	for ak, av := range charMapA {
		if av != charMapB[ak] {
			return false
		}
	}

	return true
}

func anagramSortSolution(strA string, strB string) bool {
	//`Compile` an optimized `Regexp` struct alphanumeric
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	// convert string to byte array
	sortedA := sortString(reg.ReplaceAllString(strA, ""))
	sortedB := sortString(reg.ReplaceAllString(strB, ""))

	if sortedA != sortedB {
		return false
	}
	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func charMap(str string) map[rune]int {
	r := make(map[rune]int)

	for _, c := range str {
		r[c] = r[c] + 1
	}

	return r
}
