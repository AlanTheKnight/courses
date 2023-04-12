package main

import "fmt"

func main() {
	var deg, hours, minutes int
	fmt.Scan(&deg)

	hours = deg / 30
	minutes = 2 * (deg % 30)

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
	fmt.Println()
}
