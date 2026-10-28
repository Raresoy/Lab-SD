package main

import (
	"fmt"
	"time"
)

func main() {

	ok := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		ok <- true
	}()
	<-ok
	fmt.Println("gata")
}
