package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

func letterCasePermutation(s string) []string {
	if s == "" {
		return nil
	}
	var result []string
	if len(s) > 1 {
		others := letterCasePermutation(s[1:])
	}
	r := s[0]
	if unicode.IsLetter(rune(r)) {

	} else {

	}

	for i, r := range s {
		others := letterCasePermutation(s[i:])
		s := strings.Builder{}
		if unicode.IsLetter(r) {
			result = append(result)
		}
	}
	return result
}
