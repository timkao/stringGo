package main

import (
	"fmt"
	"strconv"
)

func stringConversion() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The String is %s\n", newS)

}
