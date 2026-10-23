package main

import (
	"fmt"
)

func main() {
	sir := []int32{1, 2, 3, 4, 5, 6}
	var x int32
	fmt.Scan(&x)
	sir1 := sir[:x]
	sir2 := sir[x+1:]
	sir_nou := append(sir1, sir2...)
	fmt.Print(sir_nou)
}
