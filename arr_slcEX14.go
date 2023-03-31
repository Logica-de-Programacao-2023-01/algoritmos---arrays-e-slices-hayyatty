package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	var i, j int

	fmt.Print("Digite o primeiro indice :")
	fmt.Scan(&i)
	fmt.Println("Digite o segundo indice :")
	fmt.Scan(&j)

	aux := list[i]
	list[i] = list[j]
	list[j] = aux

	fmt.Print("O slice resultante é :", list)

}
