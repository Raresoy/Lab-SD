package main

import (
	"fmt"
	"sync"
	"time"
)

type pozitie struct {
	x, y int
}

func robot(id int, start <-chan struct{}, out chan<- pozitie, in <-chan pozitie, wg *sync.WaitGroup) {
	defer wg.Done()

	<-start
	fmt.Printf("Robot %d: Start\n", id)

	traseu := []pozitie{
		{x: id*10 + 1, y: 1},
		{x: id*10 + 2, y: 2},
		{x: id*10 + 3, y: 3},
	}
	go func() {
		for _, p := range traseu {
			time.Sleep(150 * time.Millisecond)
			out <- p
			fmt.Printf("Robot %d -> Trimis (%d, %d)\n", id, p.x, p.y)
		}
		close(out)
		fmt.Printf("Robot %d: Stop traseu\n", id)
	}()

	for p := range in {
		fmt.Printf("Robot %d <- Primit (%d, %d)\n", id, p.x, p.y)
	}

	fmt.Printf("Robot %d: Terminat\n", id)
}

func main() {
	start := make(chan struct{})
	r1ToR2 := make(chan pozitie)
	r2ToR1 := make(chan pozitie)
	var wg sync.WaitGroup
	wg.Add(2)
	go robot(1, start, r1ToR2, r2ToR1, &wg)
	go robot(2, start, r2ToR1, r1ToR2, &wg)

	fmt.Println("Observer: Start...")
	close(start)

	wg.Wait()
	fmt.Println("Observer: Gata!")
}
