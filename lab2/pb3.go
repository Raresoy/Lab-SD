package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan string)
	go func() {
		ch <- "primul"
		wg.Done()
	}()
	go func() {
		ch <- "al doilea"
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for ele := range ch {
		fmt.Println(ele)
	}
}
