package main

import "fmt"

func main() {
	fmt.Println(reverse("abcd"))
}

func reverse(str string) string {
	revstr := ""
	for _, c := range str {
		revstr = string(c) + revstr
	}
	return revstr
}
