// --- Directions
// Given a string, return true if the string is a palindrome
// or false if it is not.  Palindromes are strings that
// form the same word if it is reversed. *Do* include spaces
// and punctuation in determining if the string is a palindrome.
// --- Examples:
//   palindrome("abba") === true
//   palindrome("abcdefg") === false

package main

func palindrome(str string) bool {
	return str == reverse(str)
}

func reverse(str string) string {
	revstr := ""
	for _, c := range str {
		revstr = string(c) + revstr
	}
	return revstr
}

func palindromeArray(str string) bool {
	for index := 0; index < len(str)/2; index++ {
		if str[index] != str[len(str)-index-1] {
			return false
		}
	}
	return true
}
