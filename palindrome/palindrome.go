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
