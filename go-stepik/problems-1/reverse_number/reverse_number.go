package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 100*(n%10) + 10*((n/10)%10) + (n / 100)

	fmt.Println(ans)
}
