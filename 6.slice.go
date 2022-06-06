package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 0, 2, 1} //Creates an undefined size of array
	a = append(a, 13)         //Adds an element to the end of the slice, append(a, ) is the format
	a = append(a, 14)         //Creates a new array and stores the value automatically in the background
	fmt.Println(a)
}
