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

	// Potential Palindrome
	fmt.Println(isPalindrome("aa"))
	fmt.Println(isPalindrome("aab"))
	fmt.Println(isPalindrome("baa"))
	fmt.Println(isPalindrome("bbaa"))
	fmt.Println(isPalindrome("bbaacc"))

	// Not
	fmt.Println(isPalindrome("can"))
	fmt.Println(isPalindrome("sdlkf"))
	fmt.Println(isPalindrome("sdlkfllwlelle"))
}

func isPalindrome(x string) string {
	b := []byte(x)

	var result string
	if !palindromeFinder(b) {
		if !advancedPalindromeFinder(b) {
			result = "not "
		} else {
			result = "potentially "
		}
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

func advancedPalindromeFinder(chars []byte) bool {
	charMap := make(map[byte]int)

	for _, char := range chars {
		currentCount, ok := charMap[char]
		if !ok {
			charMap[char] = 0
			currentCount = 0
		}

		charMap[char] = currentCount + 1
	}

	numberOfOddFrequencies := 0
	for _, frequency := range charMap {
		if frequency%2 == 1 {
			numberOfOddFrequencies++
		}
	}

	return numberOfOddFrequencies < 2
}
