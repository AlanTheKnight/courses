package main

import (
	"fmt"
	"math"
	"golang.org/x/exp/constraints"
)

func MinOf[T constraints.Ordered](vars ...int) int {
    min := vars[0]

    for _, i := range vars {
        if min > i {
            min = i
        }
    }

    return min
}

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if math.Max(a, math.Max(b, c)) > b + math.Min(a, math.Min(b, c)) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
