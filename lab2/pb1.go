package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "salut"
	}()
	message := <-ch
	fmt.Println(message)
}
