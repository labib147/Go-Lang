package main

import (
	"fmt"
)

func main() {
	x := 2

	if x >= 6 {
		fmt.Println("More than 6 or equal 6 ")
	} else if x <= 2 {
		fmt.Println("Less than or equal 2")
	} else {
		fmt.Println("Between 2 and 6")
	}
}
