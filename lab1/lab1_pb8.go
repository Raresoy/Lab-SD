package main

import "fmt"

func main() {
	dict := make(map[string]int)
	result := make(map[int][]string)

	dict["Ana Popescu"] = 1
	dict["Mihai Ionescu"] = 2
	dict["Bogdan Radu"] = 1
	dict["Elena Dumitrescu"] = 3
	dict["Cristian Vasile"] = 2
	dict["Ioana Enache"] = 1

	for name, departament := range dict {
		result[departament] = append(result[departament], name)
	}
	fmt.Println(result)
}
