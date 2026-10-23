package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "primul"
	}()
	go func() {
		ch <- "al doilea"
		close(ch)
	}()

	for ele := range ch {
		fmt.Println(ele)
	}

	time.Sleep(100 * time.Millisecond)
}
