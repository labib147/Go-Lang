package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Flame", age: 22}
	fmt.Println(p)
	fmt.Println(p.name)
}
