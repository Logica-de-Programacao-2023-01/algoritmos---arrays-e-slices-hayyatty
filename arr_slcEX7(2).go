package main

import (
	"fmt"
)

func main() {
	var (
		n int
	)
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)
	fmt.Println("Digite um número")
	fmt.Scan(&n)

	for _, i := range list {
		if i == n {
			fmt.Println("Esse número está na lista")
			return
		}
	}
	list = append(list, n)
	fmt.Printf("O número %d foi adiciona a lista. \nNova Lista : %d\n ", n, list)

}
