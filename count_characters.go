package main

import (
	"fmt"
	"unicode/utf8"
)

func countCharacters() {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"

	// use len() to count "byte"
	fmt.Printf("The number of bytes in str1 is %d\n", len(str1))
	fmt.Printf("The number of bytes in str2 is %d\n", len(str2))

	// use utf8.RuneCountInString(str1) to count "Rune"
	fmt.Printf("The number of characters in str1 is %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("The number of characters in str2 is %d\n", utf8.RuneCountInString(str2))

}
