package main

import "fmt"

func main() {
	var n int
	var x int
	fmt.Scan(&n)
	fmt.Scan(&x)
	twosum := make(map[int]int)
	sir := []int{}
	result := []int{}

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		sir = append(sir, m)
	}

	for i := 0; i < n; i++ {
		_, ok := twosum[x-sir[i]]
		if !ok {
			twosum[sir[i]] = i
		} else {
			result = append(result, twosum[x-sir[i]])
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
