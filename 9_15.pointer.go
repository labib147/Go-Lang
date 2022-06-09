package main

import (
	"fmt"
)

func main() {
	i := 7
	inc(&i) // we have to use & to send the pointer
	fmt.Println(i)
	fmt.Println(&i) //& - to get the memory address of the variable
}

func inc(x *int) { // Prefixing the type with *
	*x++ // Reference the pointer to increment from. Without this it will be incrementing the memory address.
}

/*
func main() {
	i := 7
	inc(i)
	fmt.Println(i)
}

func inc(x int) {	// i is copied by value so the increment function gets a copy of i
	x++	// Does not save the value because were not returning anything
}
*/
