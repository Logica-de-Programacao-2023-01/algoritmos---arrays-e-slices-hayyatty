package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		l, c  int
		valor int
	)
	fmt.Println("Vamos inicialiar uma matriz")
	fmt.Println("Digite a quantidade de linhas")
	fmt.Scan(&l)
	fmt.Println("Digite a quantidade de colunas")
	fmt.Scan(&c)

	matriz := make([][]int, l)
	for i := 0; i < l; i++ {
		matriz[i] = make([]int, c)
	}

	for x := 0; x < 1; x++ {
		fmt.Println("Iniciando ")
		for y := 0; y < 3; y++ {
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
		}

		for i := 0; i < l; i++ {
			for j := 0; j < c; j++ {
				fmt.Printf("Digite o valor para a posição [%d][%d]:", i, j)
				fmt.Scan(&valor)
				matriz[i][j] = valor
			}
		}
		fmt.Println(matriz)
	}
}
