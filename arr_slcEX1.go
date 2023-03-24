package main

import "fmt"

func main() {
	var (
		soma int
	)
	list := [3]int{1, 2, 3}
	for _, num := range list {
		soma += num
	}
	fmt.Println("A soma da lista é ", soma)

}
