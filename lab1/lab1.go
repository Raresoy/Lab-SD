package main

import (
	"fmt"
)

func main() {
	var x int
	var div int
	fmt.Scan(&x)

	for d := 2; d < x/2; d++ {
		if x%d == 0 {
			div++
		}
	}

	if div == 0 {
		fmt.Print("prim")
	} else {
		fmt.Print("nu prim")
	}
}
