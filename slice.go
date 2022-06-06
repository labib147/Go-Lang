package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 0, 2, 1}
	a = append(a, 14)

	fmt.Println(a)
}
