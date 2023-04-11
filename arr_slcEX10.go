package main

import "fmt"

func main() {
	var (
		min int
		max int
	)
	list := []int{5, 7, 6, 1, 2, 3, 4, 7, 5, 6, 4, 8, 3}
	min = list[0]
	max = list[0]
	for _, i := range list {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}

	}
	fmt.Println(list)
	fmt.Println("O valor máximo é", max)
	fmt.Println("O valor minimo é", min)
}
