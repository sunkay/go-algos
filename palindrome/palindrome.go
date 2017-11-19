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
