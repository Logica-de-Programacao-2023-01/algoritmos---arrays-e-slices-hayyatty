package main

import "fmt"

func main() {
	var (
		valor int
	)
	matriz := [3][2]int{}
	for l := 0; l < 3; l++ {
		for c := 0; c < 2; c++ {
			fmt.Printf("Digite os números da linha %d e da coluna %d", l, c)
			fmt.Scan(&valor)
			matriz[l][c] = valor
		}
	}

	fmt.Println(matriz)
}
