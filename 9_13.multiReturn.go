package main

import (
	"errors"
	"fmt"
	"math"
)

//Go does not have exceptions
func main() {
	result, err := sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sqrt(x float64) (float64, error) { //Adding float64 for decimal values if the answer is not an integer
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers, Sorry") //errors is a new package
	}
	return math.Sqrt(x), nil
}
