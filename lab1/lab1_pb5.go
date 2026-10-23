package main

import (
	"fmt"
)

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	fmt.Print(slice)
}
