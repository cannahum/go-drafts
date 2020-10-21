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

	// Not
	fmt.Println(isPalindrome("can"))
	fmt.Println(isPalindrome("sdlkf"))
	fmt.Println(isPalindrome("sdlkfllwlelle"))
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
	switch len(chars) {
	case 0, 1:
		return true
	default:
		if chars[0] != chars[len(chars)-1] {
			return false
		}
		return palindromeFinder(chars[1 : len(chars)-1])
	}
}
