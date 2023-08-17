package main

import (
	"fmt"
)

func main() {
	var k float64
	fmt.Scan(&k)

	h := int(k / 3600)
	m := int(k)/60 - 60*h

	fmt.Printf("It is %d hours %d minutes.", h, m)
	fmt.Println()
}
