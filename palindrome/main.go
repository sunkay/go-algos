package main

func palindrome(str string) bool {
	r := reverse(str)
	if r == str {
		return true
	}

	return false
}

func reverse(str string) string {
	revstr := ""
	for _, c := range str {
		revstr = string(c) + revstr
	}
	return revstr
}
