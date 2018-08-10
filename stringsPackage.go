package main

import (
	"fmt"
	"strings"
)

func tryStringsMethods() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)

	fmt.Printf("Splitted in Slice: %v \n ", sl)

}
