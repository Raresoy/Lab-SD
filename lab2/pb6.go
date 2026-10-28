package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := make(chan struct{})
	robot1 := make(chan string)
	robot2 := make(chan string)

	go func() {
		<-start
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
		robot1 <- "Robotul1 a ajuns primul"
	}()
	go func() {
		<-start
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
		robot2 <- "Robotul2 a ajuns primul"
	}()
	close(start)

	select {
	case msg := <-robot1:
		fmt.Println(msg)
	case msg := <-robot2:
		fmt.Println(msg)
	}
}
