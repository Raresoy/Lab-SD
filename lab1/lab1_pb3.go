package main

import (
	"fmt"
)

func main() {
	sir := []int32{1, 2, 3, 4, 5, 6}
	sir_nou := make([]int32, 0, len(sir))

	for i := 0; i < len(sir); i++ {
		if sir[i]%2 == 0 {
			sir_nou = append(sir_nou, sir[i])
		}
	}
	fmt.Print(sir_nou)
}
