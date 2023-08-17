package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := (n / 100) + (n/10)%10 + n%10

	fmt.Println(ans)
}
