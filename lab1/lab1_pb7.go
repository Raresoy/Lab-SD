package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	var frecventa map[string]int = make(map[string]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		_, ok := frecventa[scanner.Text()]
		if !ok {
			frecventa[scanner.Text()] = 1
		} else {
			frecventa[scanner.Text()]++
		}
	}
	fmt.Print(frecventa)
}
