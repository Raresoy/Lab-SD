package main

import (
	"fmt"
	"sync"
)

func rutina(ch <-chan string, _wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	_wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	go rutina(ch, &wg)
	wg.Add(1)
	ch <- "salut"
	wg.Wait()
}
