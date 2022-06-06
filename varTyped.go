package main

import (
	"fmt"
)

func main() {
	/*
		:= can be used for both declaration and assignment
		= is only used for assignment
	*/
	//var x int = 5 & x:= 5 are the same things

	x := 7
	y := 8
	sum := x + y

	fmt.Println(sum)
}
