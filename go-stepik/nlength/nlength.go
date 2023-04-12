package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scan(&x)

	var divisor = int(math.Pow(10, float64(int(math.Log10(float64(x))))))

	fmt.Println(x / divisor)
}
