package main

import "fmt"

func main() {
	var (
		n float64
	)
	list := [6]float64{3.2, 4.1, 5.9, 6.7, 4.5, 9.2}
	fmt.Println("Digite um número para entrar em todas as posições")
	fmt.Scan(&n)
	for i := 0; i < len(list); i++ {
		list[i] += n
	}
	fmt.Println(list)
}
