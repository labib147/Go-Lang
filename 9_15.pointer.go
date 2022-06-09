package main

import (
	"fmt"
)

func main() {
	i := 7
	fmt.Println(i)
	fmt.Println(&i) //& - to get the memory address of the variable
}

/*
func inc(x int) {	// The increment function gets a copy of i
	x++	// Does not save the value because were not returning anything
}
*/
