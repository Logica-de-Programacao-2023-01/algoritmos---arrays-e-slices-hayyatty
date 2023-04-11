package main

import "fmt"

func main() {
	var (
		x int
		h bool
	)
	h = false
	x = 0
	list := [10]int{1, 2, 3, 5, 4, 6, 7, 8, 9, 10}
	fmt.Println(list)
	for _, i := range list {
		if i >= x {
			x = i
			h = true
		} else {
			h = false
			break
		}
	}
	if h {
		fmt.Println("A lista esta em ordem crescente")
	} else {
		fmt.Println("A lista nao esta em ordem decrescente")
	}

}
