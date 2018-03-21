// --- Directions
// Write a function that returns the number of vowels
// used in a string.  Vowels are the characters 'a', 'e'
// 'i', 'o', and 'u'.
// --- Examples
//   vowels('Hi There!') --> 3
//   vowels('Why do you ask?') --> 4
//   vowels('Why?') --> 0

package vowels

import (
	"regexp"
)

func vowels(str string) int {
	r, _ := regexp.Compile("(a|e|i|o|u|A|E|I|O|U)")

	return len(r.FindAllString(str, -1))
}
