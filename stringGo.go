package main

import "fmt"

func main() {
	str1 := "This is a raw string \n"
	str2 := `This is a raw string \n`

	fmt.Print(str1)
	fmt.Print(str2)
	fmt.Printf("str1 has %d letters\n", len(str1))
	fmt.Printf("str2 has %d letters\n", len(str2))

	concat()

	countCharacters()
}
