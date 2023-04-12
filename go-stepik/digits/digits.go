package main

import (
	"fmt"
	"math"
)

func main() {
	workArray := make([]uint8, 10)

	for i := 0; i < 10; i++ {
		var x uint8
		fmt.Scan(&x)
		workArray[i] = x
	}

	for i := 1; i <= 3; i++ {
		var a, b uint8
		fmt.Scan(&a, &b)
	}

	for _, elem := range workArray {
		fmt.Print(elem)
	}

}
