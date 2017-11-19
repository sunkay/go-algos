// --- Directions
// Given a string, return the character that is most
// commonly used in the string.
// --- Examples
// maxChar("abcccccccd") === "c"
// maxChar("apple 1231111") === "1"

package main

import "fmt"

func maxChar(str string) string {
	// map key[chars] frequency:value
	m := make(map[rune]int)

	for _, c := range str {
		m[c] = m[c] + 1
	}

	// find the highest frequency character
	prevValue := 0
	highKey := ""
	for k, v := range m {
		if v > prevValue {
			// rune to string conversion
			highKey = fmt.Sprintf("%c", k)
			prevValue = v
		}
	}

	return highKey
}
