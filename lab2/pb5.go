package main

import "fmt"

func recieve(in <-chan string, out chan<- string) {
	mesaj := <-in
	out <- mesaj
}

func send(ch chan<- string) {
	ch <- "salut"
}

func main() {
	chMesaj := make(chan string)
	chMain := make(chan string)
	go send(chMesaj)
	go recieve(chMesaj, chMain)
	mesaj := <-chMain
	fmt.Println(mesaj)
}
