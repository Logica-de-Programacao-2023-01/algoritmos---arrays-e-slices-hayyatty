package main

import (
	"fmt"
)

func main() {
	var (
		t    int
		nums int
		num  int
		soma int
	)
	fmt.Println("Qual o tamanho do slice ?")
	fmt.Scan(&t)
	list := []int{}
	for num != t {
		fmt.Println("Digite um número")
		fmt.Scan(&nums)
		num++
		list = append(list, num)
	}
	for _, som := range list {
		soma += som
	}
	fmt.Println(list)
	fmt.Println("A soma da sua lista é : ")
	fmt.Println(soma)
}
