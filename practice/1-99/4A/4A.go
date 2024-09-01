package main

import (
	"fmt"
)

func CF4A(n int) string {
	if n%2 == 0 && n > 2 {
		return "YES"
	}
	return "NO"
}

func main() {
	// You can still keep this part for direct execution
	var n int
	fmt.Scan(&n)
	fmt.Println(CF4A(n))
}
