package main

import "fmt"

func main() {
	intArr := []int{3, 2, 4, 1, 5}
	var aux int

	for i := 0; i < len(intArr)-1; i++ {
		for j := i + 1; j < len(intArr); j++ {
			if intArr[i] > intArr[j] {
				aux = intArr[i]
				intArr[i] = intArr[j]
				intArr[j] = aux

				intArr[i], intArr[j] = intArr[j], intArr[i]
			}
		}
	}
	fmt.Print(intArr)
}
