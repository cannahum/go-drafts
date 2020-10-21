package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Palindrome finder")
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("ana"))
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("kannak"))
}

func isPalindrome(x string) string {
	b := []byte(x)

	var result string
	if !palindromeFinder(b) {
		result = "not "
	}

	return fmt.Sprintf("The word '%s' is %sa palindrome", x, result)
}

func palindromeFinder(chars []byte) bool {
	if len(chars) == 0 || len(chars) == 1 {
		return true
	}

	first := chars[0]
	last := chars[len(chars)-1]
	if first != last {
		return false
	}

	rest := chars[1 : len(chars)-1]
	return palindromeFinder(rest)
}
