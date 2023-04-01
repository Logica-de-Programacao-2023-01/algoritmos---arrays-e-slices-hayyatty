package main

import "fmt"

func main() {
	var (
		encontrado bool
		ndigitado  int
	)
	list := [10]int{1, 23, 4123, 51, 43, 9543, 123, 5512, 781236}
	fmt.Println("Digite um número para verificarmos se ele está na lista")
	fmt.Scan(&ndigitado)
	for i := 0; i < len(list); i++ {
		if list[i] == ndigitado {
			fmt.Printf("O número %d está na lista", ndigitado)
			encontrado = true
			break
		}
	}
	if !encontrado {
		fmt.Println("O número não esta na lista")
	}
}
