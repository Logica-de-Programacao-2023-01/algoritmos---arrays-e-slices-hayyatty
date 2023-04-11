package main

import (
	"fmt"
)

func main() {
	var (
		ind1, ind2 int
	)
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	indx := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(list)
	fmt.Println(indx)
	fmt.Println("Qual o indice dos elementos que devem ser trocados de lugar")
	fmt.Scan(&ind1)
	fmt.Scan(&ind2)
	trocar(ind1, ind2, list)
	fmt.Println(list)
}

func trocar(ind1 int, ind2 int, list []int) {
	x := list[ind1]
	list[ind1] = list[ind2]
	list[ind2] = x
}
