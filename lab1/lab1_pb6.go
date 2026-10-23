package main

import (
	"fmt"
)

func main() {
	var dictionar map[string]int = make(map[string]int)
	sir := []string{}
	dictionar["Andrei"] = 190
	dictionar["Maria"] = 165
	dictionar["Ion"] = 182
	dictionar["Elena"] = 172
	dictionar["George"] = 178
	dictionar["Ana"] = 165

	for name, _ := range dictionar {
		sir = append(sir, name)
	}
	for i := 0; i < len(sir)-1; i++ {
		for j := i + 1; j < len(sir); j++ {
			if dictionar[sir[i]] < dictionar[sir[j]] {
				sir[i], sir[j] = sir[j], sir[i]
			}
		}
	}
	fmt.Print(sir)
}
