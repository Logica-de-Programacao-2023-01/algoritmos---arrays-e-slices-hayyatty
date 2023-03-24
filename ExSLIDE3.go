package main

import "fmt"

func main() {
	var (
		r int
	)
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)
	fmt.Println("Digite a posição para remover")
	fmt.Scan(&r)
	list = append(list[:r], list[(r+1):]...)
	fmt.Println(list)
}
